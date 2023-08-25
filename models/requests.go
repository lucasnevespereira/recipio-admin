package models

type SendInviteRequest struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	FamilyName string `json:"familyName"`
	FamilyID   string `json:"familyId"`
}
