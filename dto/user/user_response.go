package userdto

type UserResponse struct {
	ID            int    `json:"id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	IsSeller      bool   `json:"is_seller"`
	Gender        string `json:"gender"`
	Phone         string `json:"phone"`
	Photo         string `json:"photo"`
	Address       string `json:"address"`
	IsPhotoChange bool   `json:"is_photo_change"`
}