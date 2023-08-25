package models

type SendInviteRequest struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Slug     string `json:"slug"`
}
