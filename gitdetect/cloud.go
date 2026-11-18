package main

import (
	"encoding/json"
	"net/http"
)

func runOrgDetectors(org string) {
	detectGithub(org)
	detectGitlab(org)
}

func detectGitlab(org string) {
	url := "https://gitlab.com/" + org
	resp, err := http.Get(url)
	if err != nil {
		errL.Printf("'http.Get': %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		data := must(json.Marshal(NewOrg(org, Gitlab)))
		opL.Println(string(data))
	}
}

func detectGithub(org string) {
	url := "https://github.com/" + org
	resp, err := http.Get(url)
	if err != nil {
		errL.Printf("'http.Get': %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		data := must(json.Marshal(NewOrg(org, Github)))
		opL.Println(string(data))
	}
}
