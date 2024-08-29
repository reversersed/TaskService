package response

type TaskResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Due         string `json:"due_date"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
}
