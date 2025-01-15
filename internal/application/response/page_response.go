package response

type PageResponse[T any] struct {
	Page         int   `json:"page"`
	Limit        int   `json:"limit"`
	NextPage     *int  `json:"nextPage,omitempty"`
	PreviousPage *int  `json:"previousPage,omitempty"`
	TotalPages   int64 `json:"totalPages,omitempty"`
	Items        []T   `json:"items"`
}

func NewPageResponse[T any](items []T, page, pageSize int, totalRecords int64) PageResponse[T] {

	totalPages, nextPage, previousPage := calculatePagination(page, pageSize, totalRecords)

	return PageResponse[T]{
		Page:         page,
		Limit:        pageSize,
		NextPage:     nextPage,
		PreviousPage: previousPage,
		TotalPages:   totalPages,
		Items:        items,
	}
}

func calculatePagination(page, pageSize int, totalRecords int64) (totalPages int64, nextPage, previousPage *int) {
	totalPages = (totalRecords + int64(pageSize) - 1) / int64(pageSize)

	if page < int(totalPages) {
		n := page + 1
		nextPage = &n
	}
	if page > 1 {
		p := page - 1
		previousPage = &p
	}

	return totalPages, nextPage, previousPage
}
