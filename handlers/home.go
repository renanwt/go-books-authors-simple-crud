package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"Message": fmt.Sprintf("Welcome to Dialog's BookStore!"),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
