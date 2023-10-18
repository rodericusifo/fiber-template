package response

type EmployeeResponse struct {
	XID       string  `json:"xid"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Address   *string `json:"address"`
	Age       *int    `json:"age"`
	Birthday  *string `json:"birthday"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
