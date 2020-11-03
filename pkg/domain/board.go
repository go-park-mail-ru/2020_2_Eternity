package domain

type Board struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
}
