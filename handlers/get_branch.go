package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"preview/config"
	"preview/models"
)

func GetBranch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/branches/"):]
	branchID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid branch ID", http.StatusBadRequest)
		return
	}

	var branch models.Branch
	err = config.DB.QueryRow("SELECT branch_id, name, location FROM branches WHERE branch_id = ?", branchID).
		Scan(&branch.BranchID, &branch.Name, &branch.Location)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Branch not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(branch)
}
