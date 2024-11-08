package Request

type GetUserRequest struct {
	Id int `json:"id"`
}

type SignUpRequest struct {
	Fullname string `json:"fullName"`
	Username string `json:"userName" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInRequest struct {
	Username string `json:"userName" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	AssignedTo  string `json:"assignedTo"`
	CreatedAt   string `json:"createdAt"`
}
