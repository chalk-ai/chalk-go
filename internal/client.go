package internal

type OnlineQueryContext struct {
	Environment *string  `json:"environment"`
	Tags        []string `json:"tags"`
}
