package todo

type (
	//AddTodoRequest ...
	AddTodoRequest struct {
		ID string `json:"id"`
		Text string `json:"text"`
		Completed bool `json:"completed"`
	}
	//AddTodoResponse ...
	AddTodoResponse struct {
		Ok string `json:"ok"`
	}
	// GetByIDRequest ...
	GetByIDRequest struct {
		ID string `json:"id"`
	}
	// GetByIDResponse ...
	GetByIDResponse struct {
		Text string `json:"text"`
		Completed bool `json:"completed"`
	}
	//DeleteTodoRequest ...
	DeleteTodoRequest struct {
		ID string `json:"id"`
	}
	//DeleteTodoResponse ...
	DeleteTodoResponse struct {
		Deleted string `json:"deleted"`
	}
	//UpdateTodoRequest ...
	UpdateTodoRequest struct {
		ID string `json:"id"`
		Text string `json:"text"`
		Completed bool `json:"completed"`
	}
	//UpdateTodoResponse ...
	UpdateTodoResponse struct {
		Updated string `json:"updated"`
	}
)