package models

type Branch struct {
	BranchID int    `json:"branch_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
