package client

func (c *Client) OnlineQuery(request OnlineQueryParams) (OnlineQueryResult, error) {
	emptyResult := OnlineQueryResult{}

	if request.EnvironmentId == "" {
		request.EnvironmentId = c.EnvironmentId.Value
	}

	var serializedResponse onlineQueryResponseSerialized

	err := c.sendRequest(sendRequestParams{Method: "POST", URL: "v1/query/online", Body: request.serialize(), Response: &serializedResponse})
	if err != nil {
		return emptyResult, err
	}
	if len(serializedResponse.Errors) > 0 {
		serverErrors, deserializationErr := deserializeChalkErrors(serializedResponse.Errors)
		if deserializationErr != nil {
			return OnlineQueryResult{}, &ChalkErrorResponse{
				ClientError: deserializationErr,
			}
		}

		return emptyResult, &ChalkErrorResponse{ServerErrors: serverErrors}
	}

	return serializedResponse.deserialize()
}
