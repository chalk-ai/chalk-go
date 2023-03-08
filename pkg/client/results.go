package client

func (result *OnlineQueryResult) GetFeatureValue(feature string) any {
	value, found := result.values[feature]
	if !found {
		return nil
	}
	return value
}
