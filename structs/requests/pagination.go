package requests

type Pagination struct {
	Page int `form:"page"`
	PerPage int `form:"per-page"`
	Order string `form:"order"` // predefined order string like "id-desc"
}