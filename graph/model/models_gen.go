// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *User  `json:"user"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewLink struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

type NewTest struct {
	Text string `json:"text"`
}

type NewTestSubInfo struct {
	Description string `json:"description"`
	TestID      string `json:"testId"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Test struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type TestSubInfo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Test        *Test  `json:"test"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
