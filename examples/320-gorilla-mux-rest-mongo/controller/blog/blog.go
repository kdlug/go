package controller

import (
	"encoding/json"

	model "github.com/kdlug/restapi/model"

	"net/http"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts, err := model.GetAllPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return

	}

	json.NewEncoder(w).Encode(posts)
}
