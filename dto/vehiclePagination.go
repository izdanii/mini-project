package dto

type Pagination struct {
	Limit     int         `json:"limit"`
	Page      int         `json:"page"`
	TotalRows int64       `json:"total_rows"`
	Rows      interface{} `json:"rows"`
}
