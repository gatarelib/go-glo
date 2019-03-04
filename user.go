package glo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// User contains information related to a User
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var userFields = []string{"created_date", "email", "name", "username"}

// GetUser get authenticated user
// https://gloapi.gitkraken.com/v1/docs/#/Users/get_user
func (a *Glo) GetUser() (user *User, err error) {
	addr := fmt.Sprintf("%s/user", a.BaseURI)

	q := url.Values{}

	for _, field := range userFields {
		q.Add("fields", field)
	}

	data, _, err := a.jsonReq(http.MethodGet, addr, nil, q)
	if err != nil {
		return
	}

	user = &User{}
	err = json.Unmarshal(data, &user)

	return
}
