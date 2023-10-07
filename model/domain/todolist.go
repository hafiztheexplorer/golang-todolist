package domain

// ini adalah representasi dari database todolist, dengan jenis variabelnya yang agak berbeda dengan di database, dengan penulisan model json nya
type Todolist struct {
	Id              int64  `json:"id"`
	DateTime        string `json:"datetime"`
	TodolistTitle   string `json:"todolisttitle"`
	TodolistContent string `json:"todolistcontent"`
	Checked         bool   `json:"checked"`
}
