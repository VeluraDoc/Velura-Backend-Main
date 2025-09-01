package auth_dto

type AuthResponseDTO struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
