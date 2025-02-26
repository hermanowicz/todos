package typy

type Todo struct {
	TodoList  string `json:"todo_list_name,omitempty"`
	TodoTitle string `json:"todo_title"`
	TodoBody  string `json:"todo_body"`
	User      string `json:"user"`
	Status    bool   `json:"status"`
}
