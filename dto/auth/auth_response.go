package authdto

type AuthResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	IsSeller bool   `json:"is_seller"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Token    string `json:"token"`
}

type BecomeSellerResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	IsSeller bool   `json:"is_seller"`
	Token    string `json:"token"`
}