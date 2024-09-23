package taskmanagement

type Tasks struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
}
