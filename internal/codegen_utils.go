// Package internal 's codegen_utils.go holds utils that operate on
// codegen-ed structs.
package internal

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type hasManyJoin struct {
	localKey   string
	foreignKey string
}

func isHasMany(field reflect.StructField) bool {
	// TODO: Really the tag is used only for UnmarshalInto.
	//       Users don't really have to use UnmarshalInto
	//       so maybe in some places determining has-manyness
	//       by whether the type is `*[]MyHasManyFeature` is
	//       more appropriate.
	return field.Tag.Get("has_many") != ""
}

func getHasManyJoin(field reflect.StructField) hasManyJoin {
	tags := strings.Split(field.Tag.Get("has_many"), ",")
	return hasManyJoin{
		localKey:   tags[0],
		foreignKey: tags[1],
	}
}

func ResolveFeatureName(field reflect.StructField) (string, error) {
	if tag := field.Tag.Get(NameTag); tag != "" {
		return tag, nil
	}
	versioned := field.Tag.Get("versioned")
	fieldName := ChalkpySnakeCase(field.Name)
	if versioned == "true" {
		parts := strings.Split(fieldName, "_")
		nameErr := fmt.Errorf(
			"versioned feature must have a version suffix `VN` at the"+
				" end of the attribute name, but found '%s' instead",
			fieldName,
		)
		if len(parts) == 1 {
			return "", nameErr
		}
		lastPart := parts[len(parts)-1]
		if !strings.HasPrefix(lastPart, "v") {
			return "", nameErr
		}
		version := lastPart[1:]
		prefix := strings.Join(parts[:len(parts)-1], "_")
		if version == "1" {
			fieldName = prefix
		} else {
			fieldName = prefix + "@" + version
		}
	} else if strings.HasPrefix(versioned, "default(") && strings.HasSuffix(versioned, ")") {
		version := versioned[len("default(") : len(versioned)-len(")")]
		_, convertErr := strconv.Atoi(version)
		if convertErr != nil {
			return "", fmt.Errorf(
				"expected struct tag `versioned:\"default(N)\"` "+
					"where N is an integer, but found %s instead",
				versioned,
			)
		}
		if version != "1" {
			fieldName = fieldName + "@" + version
		}
	} else if versioned != "" {
		return "", fmt.Errorf(
			"expected struct tag `versioned:\"true\"` or `versioned:\"default(N)\"` "+
				"where N is an integer, but found '%s' instead",
			versioned,
		)
	}
	return fieldName, nil
}
