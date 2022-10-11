package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chrispeterjeyaraj/track-my-wealth/server/models"
)

func (h handler) GetAllUserPrefs(w http.ResponseWriter, r *http.Request) {
	var userprefs []models.UserPref

	if result := h.DB.Find(&userprefs); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userprefs)
}
