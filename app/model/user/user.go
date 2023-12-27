package user

import "encoding/json"

type User struct {
	Email      string
	Password   string
	Id         string
	IsActivate bool
}

func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
