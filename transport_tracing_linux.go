//go:build linux

package chalk

import (
	"crypto/tls"
	"net"
	"syscall"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"golang.org/x/sys/unix"
)

// TCP_INFO is a best-effort connection snapshot, not an RPC-specific RTT.
// Never read/write/close the transport's connection or change socket options.
func tcpTraceAttributes(conn net.Conn) []attribute.KeyValue {
	if tlsConn, ok := conn.(*tls.Conn); ok {
		conn = tlsConn.NetConn()
	}
	socket, ok := conn.(syscall.Conn)
	if !ok {
		return nil
	}
	raw, err := socket.SyscallConn()
	if err != nil {
		return nil
	}
	var info *unix.TCPInfo
	var sockErr error
	if err := raw.Control(func(fd uintptr) {
		info, sockErr = unix.GetsockoptTCPInfo(int(fd), unix.IPPROTO_TCP, unix.TCP_INFO)
	}); err != nil || sockErr != nil || info == nil || info.Rtt == 0 {
		return nil
	}
	return []attribute.KeyValue{
		durationAttribute("tcp_rtt_ms", time.Duration(info.Rtt)*time.Microsecond),
		durationAttribute("tcp_rtt_variation_ms", time.Duration(info.Rttvar)*time.Microsecond),
		attribute.Int64(transportAttr+"tcp_total_retransmissions", int64(info.Total_retrans)),
	}
}
