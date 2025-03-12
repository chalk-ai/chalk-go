package chalk

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"reflect"
	"time"
)

func getFqn(feature any) (fqn string, isCodegenFeature bool, err error) {
	if featureStr, ok := feature.(string); ok {
		return featureStr, false, nil
	} else if featureObj, err := UnwrapFeature(feature); err == nil {
		return featureObj.Fqn, true, nil
	} else {
		return "", false, fmt.Errorf(
			"invalid feature reference - please make sure it's a string "+
				"or a feature field from a codegen'd struct. Found: %v", feature,
		)
	}
}

// Validation for single queries
func (p *OnlineQueryParams) validateAndPopulateParamFieldsSingle() error {
	p.validatedInputs = map[string]any{}
	for k, v := range p.rawInputs {
		fqn, isCodegen, err := getFqn(k)
		if err != nil {
			return errors.Wrap(err, "validating inputs")
		}
		if isCodegen {
			p.versioned = true
		}
		p.validatedInputs[fqn] = v
	}

	p.validatedOutputs = []string{}
	for _, output := range p.rawOutputs {
		fqn, _, err := getFqn(output)
		if err != nil {
			return errors.Wrap(err, "validating outputs")
		}
		p.validatedOutputs = append(p.validatedOutputs, fqn)
	}

	p.validatedStaleness = map[string]time.Duration{}
	for k, v := range p.rawStaleness {
		fqn, _, err := getFqn(k)
		if err != nil {
			return errors.Wrap(err, "validating staleness")
		}
		p.validatedStaleness[fqn] = v
	}

	return nil
}

// Validation for bulk queries
func (p *OnlineQueryParams) validateAndPopulateParamFieldsBulk() error {
	if err := p.validateAndPopulateParamFieldsSingle(); err != nil {
		return err // Intentional no wrap
	}

	// Validate input values are lists of the same length
	referenceLen := -1
	for k, v := range p.validatedInputs {
		rVal := reflect.ValueOf(v)
		if rVal.Kind() != reflect.Slice {
			return errors.New("input values must be slices")
		}
		if referenceLen == -1 {
			referenceLen = rVal.Len()
		}
		if rVal.Len() != referenceLen {
			return errors.Newf(
				"input values must be slices of the same length, expected %d, got %d for feature '%s'",
				referenceLen, rVal.Len(), k,
			)
		}
	}
	return nil
}

func (p *OfflineQueryParams) validateAndPopulateParamFields() error {
	p.validatedInputs = map[string][]TsFeatureValue{}
	for k, v := range p.rawInputs {
		fqn, isCodegen, err := getFqn(k)
		if err != nil {
			return errors.Wrap(err, "validating inputs")
		}
		if isCodegen {
			p.versioned = true
		}
		p.validatedInputs[fqn] = v
	}

	p.validatedOutputs = []string{}
	for _, output := range p.rawOutputs {
		fqn, _, err := getFqn(output)
		if err != nil {
			return errors.Wrap(err, "validating outputs")
		}
		p.validatedOutputs = append(p.validatedOutputs, fqn)
	}

	p.validatedRequiredOutputs = []string{}
	for _, output := range p.rawRequiredOutputs {
		fqn, _, err := getFqn(output)
		if err != nil {
			return errors.Wrap(err, "validating required outputs")
		}
		p.validatedRequiredOutputs = append(p.validatedRequiredOutputs, fqn)
	}

	referenceLen := -1
	for k, v := range p.validatedInputs {
		if len(v) == 0 {
			return errors.New("input values must not be empty")
		}

		// Validate input values are the same length
		if referenceLen == -1 {
			referenceLen = len(v)
		}

		if len(v) != referenceLen {
			return errors.Newf(
				"input values must be the same length - expected %d, got %d for feature '%s'",
				referenceLen, len(v), k,
			)
		}
	}
	p.validated = true
	return nil
}
