package internal

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/ipc"
)

func CreateRequestBody(inputs arrow.Record, outputs []string) (*[]byte, error) {
	// Serialize the message
	var result bytes.Buffer
	ioWriter := bufio.NewWriter(&result)
	ipcWriter := ipc.NewWriter(ioWriter, ipc.WithSchema(inputs.Schema()))

	// Magic string header
	magicStringLen, err := ioWriter.WriteString("chal1")
	if err != nil {
		return nil, err
	}

	// Placeholder for the size of the header
	placeholder := make([]byte, 8)
	_, err = ioWriter.Write(placeholder)
	if err != nil {
		return nil, err
	}

	// Header: TODO: get the other parameters for this.
	header := map[string]any{
		"outputs":           outputs,
		"feather_body_type": "RECORD_BATCHES",
	}
	jsonBytes, err := json.Marshal(header)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize header to JSON: %w", err)
	}
	headerLength, err := ioWriter.Write(jsonBytes)
	if err != nil {
		return nil, err
	}

	// Placeholder for the size of the body
	_, err = ioWriter.Write(placeholder)
	if err != nil {
		return nil, err
	}

	// Body
	err = ipcWriter.Write(inputs)
	if err != nil {
		return nil, err
	}

	err = ipcWriter.Close()
	if err != nil {
		return nil, err
	}

	err = ioWriter.Flush()
	if err != nil {
		return nil, err
	}

	// Fill in the sizes
	resultBytes := result.Bytes()
	binary.BigEndian.PutUint64(resultBytes[magicStringLen:], uint64(headerLength))
	nonBodyLength := magicStringLen + 8 + headerLength + 8
	bodyLength := len(resultBytes) - nonBodyLength
	binary.BigEndian.PutUint64(resultBytes[magicStringLen+8+headerLength:], uint64(bodyLength))
	return &resultBytes, nil
}
