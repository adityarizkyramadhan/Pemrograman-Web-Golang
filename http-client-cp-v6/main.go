package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	client := http.Client{}

	// Hit API https://animechan.xyz/api/quotes/anime?title=naruto with method GET:
	resp, err := client.Get("https://animechan.xyz/api/quotes/anime?title=naruto")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var animechan []Animechan
	err = json.NewDecoder(resp.Body).Decode(&animechan)
	if err != nil {
		return nil, err
	}
	return animechan, nil
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(postBody)

	// Hit API https://postman-echo.com/post with method POST:
	resp, err := http.Post("https://postman-echo.com/post", "application/json", requestBody)
	if err != nil {
		return Postman{}, err
	}
	defer resp.Body.Close()

	var postman Postman
	err = json.NewDecoder(resp.Body).Decode(&postman)
	if err != nil {
		return Postman{}, err
	}
	return postman, nil
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
