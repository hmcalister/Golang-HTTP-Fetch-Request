package main

import (
	"encoding/json"
	"net/http"
)

type Content struct {
	Name string
	Age  int
}

func GetContent(w http.ResponseWriter, req *http.Request) {
	content := Content{
		"Hayden",
		22,
	}

	// json_content, _ := json.Marshal(&content)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(content)

	// fmt.Print(string(json_content))
	// fmt.Fprint(w, json_content)
}

func main() {

	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	http.HandleFunc("/content", GetContent)

	http.ListenAndServe(":8080", nil)
}
