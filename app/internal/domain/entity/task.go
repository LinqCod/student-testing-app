package entity

type Task struct {
	Id          int64        `json:"id"`
	CategoryId  int64        `json:"category_id"`
	Description string       `json:"description"`
	Answers     []TaskAnswer `json:"answers"`
}
