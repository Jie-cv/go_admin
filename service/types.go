package service

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResonse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type QueryRequest struct {
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	Keyword string `json:"keyword" form:"keyword"`
}

type GetUserListRequest struct {
	*QueryRequest
}

type GetUserListRsponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Remarks   string `json:"remarks"`
	Email     string `json:"email"`
	Sex       string `json:"sex"`
}

type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Remarks  string `json:"remarks"`
}

type UpdateUserRequest struct {
	ID uint `json:"id"`
	AddUserRequest
}

type DeleteUserRequest struct {
	ID uint `json:"id"`
}

type RoleListRequest struct {
	*QueryRequest
}

type RoleListResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Is_Admin  int8   `json:"is_admin"`
	Remarks   string `json:"remarks"`
	Sort      int64  `json:"sort"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AddRoleRequest struct {
	Name     string `json:"name"`
	Is_Admin int8   `json:"is_admin"`
	Remarks  string `json:"remarks"`
	Sort     int64  `json:"sort"`
}

type UpdateRoleRequest struct {
	ID uint `json:"id"`
	AddRoleRequest
}

type DeleteRoleRequest struct {
	ID uint `json:"id"`
}
