package response

type LoginResponse struct {
	Success bool `json:"success"`
	UserID  uint `json:"userId"`
}