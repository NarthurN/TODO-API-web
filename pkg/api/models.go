package api

type Task struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type Response struct {
	ID    int64  `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}
