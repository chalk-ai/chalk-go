package proto

import (
	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
	"github.com/cockroachdb/errors"
	"google.golang.org/protobuf/runtime/protoimpl"
	"math/big"
	"strings"
)

type DecimalValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Precision int32  `protobuf:"varint,2,opt,name=precision,proto3" json:"precision,omitempty"`
	Scale     int32  `protobuf:"varint,3,opt,name=scale,proto3" json:"scale,omitempty"`
}

func convertDecimalToProto(value *big.Float) (*arrowv1.DecimalValue, error) {
	// Convert the big.Float to a string with no exponent part.
	valueStr := value.Text('f', -1) // Using 'f' to avoid scientific notation

	// Split the string on the decimal point to separate the whole number and the fractional part.
	parts := strings.Split(valueStr, ".")
	scale := int32(0)
	var whole string = parts[0]
	var fractional string

	if len(parts) > 1 {
		fractional = parts[1]
		scale = int32(len(fractional))
	}

	// Concatenate whole and fractional parts to form an integer string.
	intStr := whole + fractional

	// Convert the concatenated string back to big.Int.
	bigIntVal, ok := new(big.Int).SetString(intStr, 10)
	if !ok {
		return nil, errors.New("failed to convert big.Float significant digits to big.Int")
	}

	// Precision is the total number of digits.
	precision := int32(len(whole) + len(fractional))

	// Create a DecimalValue protobuf message.
	return &arrowv1.DecimalValue{
		Value:     bigIntVal.Bytes(), // Convert big.Int to a slice of bytes
		Precision: precision,
		Scale:     scale,
	}, nil
}
