package page

type PageInfo struct {
	CurrentPage int64 `form:"currentPage"`
	PageSize    int64 `form:"pageSize"`
}
