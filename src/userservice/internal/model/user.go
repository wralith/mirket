package model

type User struct {
	Id    uint32 `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Bio   string `json:"bio,omitempty"`
}
