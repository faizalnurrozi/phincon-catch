package view_models

type PaginationVm struct {
	Pagination DetailPaginationVm `json:"pagination"`
}

type DetailPaginationVm struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
}

func NewPaginationVm() PaginationVm {
	return PaginationVm{}
}

func (vm PaginationVm) Build(detailPagination DetailPaginationVm) PaginationVm {
	return PaginationVm{Pagination: detailPagination}
}
