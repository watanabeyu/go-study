package main

import (
	fb "github.com/huandu/facebook"

	"database/sql"
	"fmt"
)

type fbResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	Location loc    `json:"location"`
}

type loc struct {
	Id   sql.NullString `json:"id"`
	Name sql.NullString `json:"name"`
}

func mapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//connect
	var globalApp = fb.New("AppId", "AppSecret")
	session, _ := globalApp.SessionFromSignedRequest("facebook signed request")

	err := session.Validate()

	if err != nil {
		panic(err)
	}

	//get me
	response, _ := session.Get("/me", fb.Params{
		"fields": "id,name,email,birthday,gender,location",
		"locale": "ja_JP",
	})

	//response to struct
	var fbresponse fbResponse
	mapToStruct(response, &fbresponse)

	fmt.Println(fbresponse)
}
