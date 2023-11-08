package struct_test

//	@Summary		struct
//	@Description	struct for getting the credentials
//	@Tags			credentials
type AdminAcc struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}
