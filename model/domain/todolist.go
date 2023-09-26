package domain

type Todolist struct {
	Id              int64  `json:"id"`
	DateTime        string `json:"datetime"`
	TodolistTitle   string `json:"todolisttitle"`
	TodolistContent string `json:"todolistcontent"`
	Checked         bool   `json:"checked"`
}
