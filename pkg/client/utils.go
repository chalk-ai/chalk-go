package client

func stringPointerOrNil(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
