package my_project

import (
	"encoding/json"
	"net/http"
	"time"

type TimeResponse struct {
	Time string `json:"time"`
}

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		response := TimeResponse{Time: currentTime}
	}

}