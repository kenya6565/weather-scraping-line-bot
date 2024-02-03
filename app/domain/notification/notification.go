package domain

type LineEvent struct {
	Events []struct {
		Type       string `json:"type"`
		ReplyToken string `json:"replyToken"`
		Source     struct {
			UserID string `json:"userId"`
		} `json:"source"`
		Message struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}
