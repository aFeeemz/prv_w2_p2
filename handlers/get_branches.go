package handlers

import (
	"encoding/json"
	"net/http"

	"preview/config"
	"preview/models"
)

func GetBranches(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query("SELECT branch_id, name, location FROM branches")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var branches []models.Branch
	for rows.Next() {
		var branch models.Branch
		if err := rows.Scan(&branch.BranchID, &branch.Name, &branch.Location); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		branches = append(branches, branch)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(branches)
}
