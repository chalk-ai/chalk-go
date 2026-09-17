//go:build !linux

package chalk

import (
	"net"

	"go.opentelemetry.io/otel/attribute"
)

func tcpTraceAttributes(net.Conn) []attribute.KeyValue { return nil }
