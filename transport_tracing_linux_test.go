//go:build linux

package chalk

import (
	"crypto/tls"
	"io"
	"net"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestTransportTracingTCPInfo(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	assert.NoError(t, err)
	defer listener.Close()
	done := make(chan error, 1)
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			done <- err
			return
		}
		defer conn.Close()
		_, err = io.CopyN(conn, conn, 1)
		done <- err
	}()
	conn, err := net.Dial("tcp", listener.Addr().String())
	assert.NoError(t, err)
	defer conn.Close()
	_, err = conn.Write([]byte{1})
	assert.NoError(t, err)
	_, err = io.ReadFull(conn, make([]byte, 1))
	assert.NoError(t, err)
	assert.NoError(t, <-done)
	for _, socket := range []net.Conn{conn, tls.Client(conn, &tls.Config{ServerName: "localhost"})} {
		attrs := tcpTraceAttributes(socket)
		assert.Len(t, attrs, 3)
		assert.Equal(t, transportAttr+"tcp_rtt_ms", string(attrs[0].Key))
		assert.Positive(t, attrs[0].Value.AsFloat64())
	}
	assert.NoError(t, conn.Close())
	assert.Empty(t, tcpTraceAttributes(conn), "closed sockets must not affect the query")
	pipe, peer := net.Pipe()
	defer pipe.Close()
	defer peer.Close()
	assert.Empty(t, tcpTraceAttributes(pipe), "unsupported wrappers must be harmless")
}
