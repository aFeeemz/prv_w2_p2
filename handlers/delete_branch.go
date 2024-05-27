package handlers

import (
	"net/http"
	"strconv"

	"preview/config"
)

func DeleteBranch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/branches/delete/"):]
	branchID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid branch ID", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("DELETE FROM branches WHERE branch_id = ?", branchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
