package utils

import (
	"net/url"
	"strconv"
)

// GetPaginationParams - gets params from query or returns default
func GetPaginationParams(q url.Values) (limit int, offset int) {
	maxLimit := 200
	page := 1
	if pageStr, ok := q["page"]; ok && len(pageStr) > 0 {
		p, err := strconv.Atoi(pageStr[0])
		if err == nil && p > 0 {
			page = p
		}
	}

	limit = 20
	if limitStr, ok := q["limit"]; ok && len(limitStr) > 0 {
		l, err := strconv.Atoi(limitStr[0])
		if err == nil && l > 0 && l <= maxLimit {
			limit = l
		}
	}

	if page > 1 {
		offset = (page - 1) * limit
	}

	return
}
