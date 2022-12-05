package entity

type TaskCategory struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}
