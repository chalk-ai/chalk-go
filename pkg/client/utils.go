package client

func getStringifiedInputs(inputs map[string]interface{}) map[string]string {
	res := make(map[string]string)
	for fqn, value := range inputs {
		valueStr := value.(string)
		res[fqn] = valueStr
	}

	return res
}
