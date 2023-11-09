package struct_test

// @Summary		struct
// @Description	struct for getting the credentials
// @Tags			credentials
type AdminAcc struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Password   string `json:"password"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}
