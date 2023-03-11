package chalk

func (result *OnlineQueryResult) GetFeature(feature string) *FeatureResult {
	featureResult, found := result.features[feature]
	if !found {
		return nil
	}
	return &featureResult
}

func (result *OnlineQueryResult) GetFeatureValue(feature string) any {
	featureResult := result.GetFeature(feature)
	if featureResult == nil {
		return nil
	}
	return featureResult.Value
}
