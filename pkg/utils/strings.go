package utils

func StrPtrOrNil(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
