package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DefaultResponse struct {
	IsOK bool        `json:"is_ok"`
	Data interface{} `json:"data"`
}

func SendOK(w http.ResponseWriter, code int, p interface{}) {
	w.WriteHeader(code)

	//Todo: wrap response into html page (as tasked)
	err := json.NewEncoder(w).Encode(DefaultResponse{
		IsOK: true,
		Data: p,
	})

	if err != nil {
		fmt.Println("Handler error")
	}
}

