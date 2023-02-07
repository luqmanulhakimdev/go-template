package services

import "go-template/helpers/str"

var (
	paginationdefaultLimit    = 10
	paginationmaxLimit        = 50
	paginationdefaultSort     = "asc"
	paginationsortWhitelist   = []string{"asc", "desc"}
	paginationdefaultLastPage = 0
)

type Pagination struct {
	CurrentPage int `json:"current_page" example:"1"`
	LastPage    int `json:"last_page" example:"1"`
	Total       int `json:"total" example:"1"`
	PerPage     int `json:"per_page" example:"10"`
}

func SetPaginationParameter(page, limit int, orderBy, sort string, orderByWhiteLists, orderByStringWhiteLists []string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > paginationmaxLimit {
		limit = paginationdefaultLimit
	}

	orderBy = checkWhiteList(orderBy, orderByWhiteLists)
	if str.Contains(orderByStringWhiteLists, orderBy) {
		orderBy = `LOWER(` + orderBy + `)`
	}

	if !str.Contains(paginationsortWhitelist, sort) {
		sort = paginationdefaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, orderBy, sort
}

func checkWhiteList(orderBy string, whiteLists []string) string {
	for _, whiteList := range whiteLists {
		if orderBy == whiteList {
			return orderBy
		}
	}

	if len(whiteLists) == 0 {
		return "def.updated_at"
	}

	return whiteLists[0]
}

func SetPaginationResponse(page, limit, total int) (pagination Pagination) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = paginationdefaultLastPage
	}

	pagination = Pagination{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	}

	return pagination
}
