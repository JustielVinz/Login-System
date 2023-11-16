package struct_test

// @Summary		struct
// @Description	struct for getting the credentials
// @Tags			credentials
type AdminAcc struct {
	Name       string `json:"name"`
	Department string `json:"department"`
}
type Student struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	StudentID     int    `json:"student_id"`
	Department    string `json:"department"`
	Miscellaneous string `json:"miscellaneous"`
	Payment       int    `json:"payment_method"`
	Amount        int    `json:"amount"`
}
