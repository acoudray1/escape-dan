package models

import (
    "net/http"
)

type User struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Phone string `json:"phone"`
	Mail string `json:"mail"`
}

type UsersList struct {
    Users []User `"json:users"`
}

func (i *User) Bind(r *http.Request) error {
    // if i.Name == "" {
    //     return fmt.Errorf("name is a required field")
    // }
    return nil
}

func (*UsersList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
