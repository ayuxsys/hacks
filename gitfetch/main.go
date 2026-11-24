package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	token := os.Getenv("GH_TOKEN")
	if token == "" {
		os.Exit(1)
	}
	itms := gitFetch(os.Args[1], token)
	for _, i := range itms {
		if i != "" {
			fmt.Println(i)
		}
	}
}

func gitFetch(url, token string) []string {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer: %s", token))
	req.Header.Add("Accept", "application/vnd.github.raw")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(body), "\n")
}
