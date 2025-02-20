package dto

type ForgetPasswordDTO struct {
	UserID      uint64 `json:"user_id"`
	NewPassword string `json:"new_password"`
}
