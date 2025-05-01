package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
)

type User struct {
	Role       string `json:"role"`
	ID         string `json:"id"`
	Experience int    `json:"experience"`
	Remote     bool   `json:"remote"`
	User       struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func getResources(path string) []map[string]any {
	fullURL, err := url.Parse("https://api.boot.dev")
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil
	}

	fullURL.Path, err = url.JoinPath(fullURL.Path, path)
	if err != nil {
		fmt.Println("Error joining path:", err)
		return nil
	}

	res, err := http.Get(fullURL.String())
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	defer res.Body.Close()

	var resources []map[string]any
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resources)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}

	return resources
}

func logUser(user User) {
	fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n",
		user.User.Name, user.Role, user.Experience, user.Remote)
}

func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n",
			user.User.Name, user.Role, user.Experience, user.Remote)
	}
}

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL, err := url.Parse(baseURL + "/" + id)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return User{}, err
	}

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return User{}, err
	}

	req, err := http.NewRequest(http.MethodPut, fullURL.String(), bytes.NewBuffer(b))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var updatedUser User
	err = json.NewDecoder(resp.Body).Decode(&updatedUser)
	if err != nil {
		return User{}, err
	}

	return updatedUser, nil
}

func getUsers(rawUrl string) ([]User, error) {
	fullURL, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	q := fullURL.Query()
	q.Add("sort", "experience")
	fullURL.RawQuery = q.Encode()

	res, err := http.Get(fullURL.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var users []User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL, err := url.Parse(baseURL + "/" + id)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return User{}, err
	}

	req, err := http.NewRequest(http.MethodGet, fullURL.String(), nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("X-API-Key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func deleteUser(baseURL, id, apiKey string) error {
	fullURL, err := url.Parse(baseURL + "/" + id)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, fullURL.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-API-Key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func main() {
	baseURL := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"

	users, err := getUsers(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	logUsers(users)
	userId := "2f8282cb-e2f9-496f-b144-c0aa4ced56db"
	baseURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	apiKey := generateKey()

	userData, err := getUserById(baseURL, userId, apiKey)
	if err != nil {
		fmt.Println(err)
	}
	logUser(userData)

	fmt.Printf("Updating user with id: %s\n", userData.ID)
	userData.Role = "Senior Backend Developer"
	userData.Experience = 7
	userData.Remote = true
	userData.User.Name = "Allan"

	updatedUser, err := updateUser(baseURL, userId, apiKey, userData)
	if err != nil {
		fmt.Println(err)
		return
	}
	logUser(updatedUser)
}

func generateKey() string {
	const characters = "ABCDEF0123456789"
	result := ""
	rand.New(rand.NewSource(0))
	for i := 0; i < 16; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}
