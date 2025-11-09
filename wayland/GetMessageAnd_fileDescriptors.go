package wayland

import (
	"errors"
	"io"
	"net"
	"syscall"
	"time"
)

func GetMessageAndFileDescriptors(conn *net.UnixConn, buf []byte) (n int, fds []int, shouldContinue bool, err error) {
	const (
		timeout = 10 * time.Millisecond
		maxFDs  = 255
		fdSize  = 4 // bytes per C int in SCM_RIGHTS payload
	)

	// Prepare OOB buffer big enough for up to maxFDs.
	oob := make([]byte, syscall.CmsgSpace(maxFDs*fdSize))

	// Apply a short read deadline to emulate select timeout.
	if err := conn.SetReadDeadline(time.Now().Add(timeout)); err != nil {
		return 0, nil, false, err
	}
	defer conn.SetReadDeadline(time.Time{})

	n, oobn, flags, _, rerr := conn.ReadMsgUnix(buf, oob)

	if ne, ok := rerr.(net.Error); ok && ne.Timeout() {
		return 0, nil, true, nil
	}
	if rerr != nil {
		if errors.Is(rerr, io.EOF) {
			return n, nil, false, nil
		}
		return n, nil, false, rerr
	}
	if n == 0 {
		return 0, nil, false, nil
	}

	if flags&syscall.MSG_TRUNC != 0 {
		return n, nil, false, errors.New("payload truncated (MSG_TRUNC)")
	}
	if flags&syscall.MSG_CTRUNC != 0 {
		return n, nil, false, errors.New("control data truncated (MSG_CTRUNC)")
	}

	if oobn > 0 {
		cmsgs, perr := syscall.ParseSocketControlMessage(oob[:oobn])
		if perr != nil {
			return n, nil, false, perr
		}
		for _, cmsg := range cmsgs {
			rights, rerr := syscall.ParseUnixRights(&cmsg)
			if rerr == nil && len(rights) > 0 {
				fds = append(fds, rights...)
				if len(fds) >= maxFDs {
					fds = fds[:maxFDs]
					break
				}
			}
		}
	}

	return n, fds, true, nil
}
