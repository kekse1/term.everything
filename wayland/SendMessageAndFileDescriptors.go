package wayland

import (
	"io"
	"net"
	"syscall"
)

func SendMessageAndFileDescriptors(conn *net.UnixConn, buf []byte, fds []int) (int, bool, error) {
	var oob []byte
	if len(fds) > 0 {
		oob = syscall.UnixRights(fds...)
	}

	n, oobn, err := conn.WriteMsgUnix(buf, oob, nil)
	if err != nil {
		return 0, false, err
	}
	if len(oob) > 0 && oobn < len(oob) {
		return n, false, io.ErrShortWrite
	}

	return n, true, nil
}
