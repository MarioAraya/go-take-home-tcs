package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 2 * time.Second}

func FetchUser(id int) (*User, error) {
	url := fmt.Sprintf("https://dummyjson.com/users/%d", id)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func FetchTodos(id int) (*TodosResponse, error) {
	url := fmt.Sprintf("https://dummyjson.com/todos/user/%d", id)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var todos TodosResponse
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		return nil, err
	}
	return &todos, nil
}
