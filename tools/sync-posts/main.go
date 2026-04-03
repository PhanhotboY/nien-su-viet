package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Post struct {
	Id          string  `json:"id"`                  // Primary key
	Title       string  `json:"title"`               // Post title
	Slug        string  `json:"slug"`                // URL slug
	Content     string  `json:"content"`             // Post content
	Summary     *string `json:"summary,omitempty"`   // Post summary
	Thumbnail   *string `json:"thumbnail,omitempty"` // Thumbnail URL
	AuthorId    string  `json:"authorId"`            // Foreign key to author
	CategoryId  string
	Views       int    `json:"views"`                 // Number of views
	Likes       int    `json:"likes"`                 // Number of likes
	Published   bool   `json:"published"`             // Publication status
	PublishedAt string `json:"publishedAt,omitempty"` // Publication timestamp
	CreatedAt   string `json:"createdAt"`             // Creation timestamp
	UpdatedAt   string `json:"updatedAt"`             // Last update timestamp
}

func main() {
	postsByte, err := os.ReadFile("../../posts.json")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer os.Exit(0)

	var posts []Post
	err = json.Unmarshal(postsByte, &posts)
	if err != nil {
		log.Fatalf("error unmarshaling posts: %v", err)
	}

	cookie, err := login()
	if err != nil {
		log.Fatalf("error logging in: %v", err)
	}

	for i, p := range posts {
		createPost(i, p, cookie)
	}

}

func createPost(i int, p Post, cookie string) {
	fmt.Printf("Start creating %dth post\n", i)

	body, err := json.Marshal(p)
	if err != nil {
		fmt.Print(fmt.Errorf("error marshaling create %dth post body: %v", i, err))
		return
	}
	req, err := http.NewRequest("POST", os.Getenv("API_GATEWAY_ENDPOINT")+"/api/v1/posts", bytes.NewBuffer(body))
	if err != nil {
		fmt.Print(fmt.Errorf("error creating %dth post: %v", i, err))
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("cookie", cookie)

	client := http.Client{}
	res, err := client.Do(req)
	if res.StatusCode != 201 || err != nil {
		fmt.Print(fmt.Errorf("failed to create %dth post: %v", i, err))
	}
	defer res.Body.Close()

	fmt.Printf("\ncreated post: %s", p.Title)
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login() (cookie string, err error) {
	credentials := LoginCredentials{
		Email:    os.Getenv("EDITOR_EMAIL"),
		Password: os.Getenv("EDITOR_PASSWORD"),
	}

	body, err := json.Marshal(credentials)
	if err != nil {
		return "", fmt.Errorf("error marshaling login credentials: %v", err)
	}
	req, err := http.NewRequest("POST", os.Getenv("API_GATEWAY_ENDPOINT")+"/api/v1/auth/sign-in/email", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("origin", os.Getenv("CLIENT_ORIGIN"))

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	cookies := make([]string, 0)
	for _, c := range res.Header.Values("set-cookie") {
		cookies = append(cookies, strings.Split(c, ";")[0])
	}

	return strings.Join(cookies, "; "), nil
}
