package model

type Note struct {
	Title      string   `json:"title"`
	TitleDesc  string   `json:"title_desc"`
	Content    string   `json:"content"`
	Labels     []string `json:"labels"`
	Author     string   `json:"author"`
	CreateTime string   `json:"create_time"`
	UpdateTime string   `json:"update_time"`
}

type NoteLabel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
