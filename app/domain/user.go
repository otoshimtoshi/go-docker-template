package domain

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	Birthday  time.Time
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
	ImageUrl  string `json:"image_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) ToUserResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Birthday:  u.Birthday.Format("2006-01-02"),
		ImageUrl:  u.ImageUrl,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
}
