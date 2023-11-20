package models

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ImageUrl  string `json:"imageUrl"`
	Guest     bool   `json:"guest"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	IsDeleted string `json:"isDeleted"`
}
