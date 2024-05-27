package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"preview/config"
	"preview/models"
)

func UpdateBranch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/branches/update/"):]
	branchID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid branch ID", http.StatusBadRequest)
		return
	}

	var branch models.Branch
	if err := json.NewDecoder(r.Body).Decode(&branch); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("UPDATE branches SET name = ?, location = ? WHERE branch_id = ?", branch.Name, branch.Location, branchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
