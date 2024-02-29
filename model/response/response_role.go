package response

type Role struct {
	Role_id   int    `json:"role_id"`
	Role_name string `json:"role_name"`

	User []Response_user `json:"user"`
}
