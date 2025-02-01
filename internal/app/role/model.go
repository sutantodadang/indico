package role

type RegisterRoleRequest struct {
	UniqueName string `json:"unique_name" binding:"required"`
	Name       string `json:"name" binding:"required"`
}

type Role struct {
	UserRoleID string `json:"user_role_id"`
	UniqueName string `json:"unique_name"`
	Name       string `json:"name"`
	Status     bool   `json:"status"`
}
