package common

var (
	ORDER_ASCENDING  = "ASC"
	ORDER_DESCENDING = "DESC"
)

type CommonQueryParams struct {
	Skip    uint16
	Limit   uint16
	OrderBy string
	Order   string
}

type PaginationResponse struct {
	Total uint32
}
