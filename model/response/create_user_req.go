package response

type Create_req struct {
	User_name string `validate:"required,max=100,min=1" json:"user_name"`
	User_pass string `validate:"required,max=100,min=1" json:"user_pass"`
	Role_id   int    `validate:"required" json:"role_id"`
}
