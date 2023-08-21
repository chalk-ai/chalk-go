package internal

import (
	"bufio"
	"bytes"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/ipc"
)

func CreateRequestBody(inputs arrow.Record) (*[]byte, error) {
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
	err = ioWriter.WriteByte(0)
	if err != nil {
		return nil, err
	}

	// Header: TODO: get the other parameters for this.
	headerLength, err := ioWriter.WriteString(`{"output": [...]}`)
	if err != nil {
		return nil, err
	}

	// Placeholder for the size of the body
	err = ioWriter.WriteByte(0)
	if err != nil {
		return nil, err
	}

	// Body
	err = ipcWriter.Write(inputs)
	if err != nil {
		return nil, err
	}

	// Fill in the sizes
	err = ipcWriter.Close()
	if err != nil {
		return nil, err
	}

	err = ioWriter.Flush()
	if err != nil {
		return nil, err
	}
	resultBytes := result.Bytes()
	resultBytes[magicStringLen] = uint8(headerLength)
	resultBytes[magicStringLen+1+headerLength] = uint8(
		len(resultBytes) - magicStringLen - 1 - headerLength,
	)
	return &resultBytes, nil
}
