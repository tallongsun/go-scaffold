package dto

import "time"

type UserCreateReq struct {
	Name string `json:"name"`
}

type UserUpdateReq struct {
	Name string `json:"name"`
}

type UserRes struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
