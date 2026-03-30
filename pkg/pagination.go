package pkg

import (
	"math"
	"net/http"
	"strconv"
)

type PaginationQuery struct {
	Page   int
	Limit  int
	Search string
	Filter string
	Sort   string
}

func GetPaginationParams(r *http.Request) PaginationQuery {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(q.Get("limit"))
	if limit <= 0 {
		limit = 10
	}

	return PaginationQuery{
		Page:   page,
		Limit:  limit,
		Search: q.Get("search"),
		Filter: q.Get("filter"),
		Sort:   q.Get("sort"),
	}
}

func CalculateTotalPages(totalData, limit int) int {
	if limit <= 0 {
		return 0
	}
	return int(math.Ceil(float64(totalData) / float64(limit)))
}
