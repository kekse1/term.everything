package wayland

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

type SocketListener struct {
	WaylandDisplayName string
	SocketPath         string
	listener           *net.UnixListener

	mu          sync.Mutex
	connections map[*net.UnixConn]struct{}

	Connections      chan *net.UnixConn
	CloseConnections chan *net.UnixConn

	closeOnce sync.Once
	closed    chan struct{}
}

type HasDisplayName interface {
	WaylandDisplayName() string
}

func NewWaylandSocketListener(args HasDisplayName) (*SocketListener, error) {
	displayName := getWaylandDisplayName(args)
	socketPath := getSocketPathFromName(displayName)

	ln, fd, err := listenToWaylandSocket(displayName)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on wayland socket: %w", err)
	}
	_ = fd

	w := &SocketListener{
		WaylandDisplayName: displayName,
		SocketPath:         socketPath,
		listener:           ln,
		connections:        make(map[*net.UnixConn]struct{}),
		Connections:        make(chan *net.UnixConn),
		CloseConnections:   make(chan *net.UnixConn),
		closed:             make(chan struct{}),
	}

	onExit(func() {
		_ = w.Close()
	})

	go w.closeConnLoop()

	return w, nil
}

func (w *SocketListener) closeConnLoop() {
	for {
		select {
		case <-w.closed:
			return
		case c := <-w.CloseConnections:
			if c == nil {
				continue
			}
			_ = c.Close()
			w.removeConnection(c)
		}
	}
}

func (w *SocketListener) MainLoop() error {
	defer w.Close()

	for {
		_ = w.listener.SetDeadline(time.Now().Add(2 * time.Second))
		conn, err := w.listener.AcceptUnix()
		if ne, ok := err.(net.Error); ok && ne.Timeout() {
			select {
			case <-w.closed:
				return nil
			default:
				continue
			}
		}
		if errors.Is(err, net.ErrClosed) || errors.Is(err, os.ErrClosed) {
			return nil
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "accept error: %v\n", err)
			continue
		}

		w.addConnection(conn)

		// Deliver the connection to consumers.
		select {
		case w.Connections <- conn:
		case <-w.closed:
			_ = conn.Close()
			return nil
		}
	}
}

func (w *SocketListener) Close() error {
	var firstErr error
	w.closeOnce.Do(func() {
		close(w.closed)

		if w.listener != nil {
			if err := w.listener.Close(); err != nil && firstErr == nil {
				firstErr = err
			}
		}

		w.mu.Lock()
		for c := range w.connections {
			_ = c.Close()
		}
		w.connections = make(map[*net.UnixConn]struct{})
		w.mu.Unlock()

		if err := removeFileIfExists(w.SocketPath); err != nil && firstErr == nil {
			firstErr = err
		}
	})
	return firstErr
}

func (w *SocketListener) addConnection(c *net.UnixConn) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.connections[c] = struct{}{}
}

func (w *SocketListener) removeConnection(c *net.UnixConn) {
	w.mu.Lock()
	defer w.mu.Unlock()
	delete(w.connections, c)
}

func onExit(callback func()) {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
	)
	go func() {
		<-ch
		callback()
		os.Exit(0)
	}()
}

func getWaylandDisplayName(args HasDisplayName) string {
	if args.WaylandDisplayName() != "" {
		return args.WaylandDisplayName()
	}
	if v := os.Getenv("WAYLAND_DISPLAY_NAME"); v != "" {
		return v
	}

	for i := 2; i < 1000; i++ {
		name := fmt.Sprintf("wayland-%d", i)
		path := getSocketPathFromName(name)
		if _, err := os.Stat(path); err == nil {
			continue
		} else if os.IsNotExist(err) {
			return name
		} else {
			continue
		}
	}
	fmt.Fprintf(os.Stderr, "Failed to find an open wayland socket name. Pass one with --wayland-display-name <name>\n")
	os.Exit(1)
	return ""
}

func getSocketPathFromName(socketName string) string {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		runtimeDir = "/tmp"
	}
	return filepath.Join(runtimeDir, socketName)
}

func removeFileIfExists(p string) error {
	if p == "" {
		return fmt.Errorf("empty path")
	}
	if _, err := os.Lstat(p); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return os.Remove(p)
}

func listenToWaylandSocket(socketName string) (*net.UnixListener, int, error) {
	socketPath := getSocketPathFromName(socketName)

	if err := removeFileIfExists(socketPath); err != nil {
		return nil, -1, fmt.Errorf("remove existing socket: %w", err)
	}

	addr := &net.UnixAddr{
		Name: socketPath,
		Net:  "unix",
	}
	ln, err := net.ListenUnix("unix", addr)
	if err != nil {
		return nil, -1, fmt.Errorf("listen unix: %w", err)
	}

	file, err := ln.File()
	if err != nil {
		_ = ln.Close()
		return nil, -1, fmt.Errorf("get listener file: %w", err)
	}
	fd := int(file.Fd())
	_ = file.Close()

	return ln, fd, nil
}
