package dto

type UserDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type InsertUserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
