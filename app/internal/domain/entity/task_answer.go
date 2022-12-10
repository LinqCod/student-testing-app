package entity

type TaskAnswer struct {
	Id      int64  `json:"id"`
	Text    string `json:"text"`
	IsRight bool   `json:"is_right"`
}
