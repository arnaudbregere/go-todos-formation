package serverweb

import (
	"encoding/json"
	"fmt"
	"formation/api"
	"net/http"
)

func ListTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("create", len(api.List()))
	json.NewEncoder(w).Encode(api.List())
	fmt.Println( api.List())
}