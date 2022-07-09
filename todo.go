package rest_api_auth

type ToDoList struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	ID     uint64
	UserID uint64
	ListID uint64
}

type ToDoItem struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        string `json:"done"`
}

type ListItem struct {
	ID     uint64
	ListID uint64
	ItemID uint64
}
