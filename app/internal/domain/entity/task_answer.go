package entity

type TaskAnswer struct {
	Id      int64  `json:"id"`
	TaskId  int64  `json:"task_id"`
	Text    string `json:"text"`
	IsRight bool   `json:"is_right"`
}
