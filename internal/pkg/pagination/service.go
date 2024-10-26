package pagination

import (
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ramasuryananda/dummy-cv/internal/constant"
)

func GetOffset(page int, limit int) int {
	if limit < 0 {
		limit = constant.LIMIT
	}

	if page > 0 {
		return ((limit * page) - limit)
	}
	return 0
}

func GetPrevPage(c echo.Context, page int) string {
	if page > 1 {
		originalParams := c.QueryParams()

		// Create a new set of query parameters
		newParams := url.Values{}

		// Copy the original query parameters to the new parameters
		for key, values := range originalParams {
			for _, value := range values {
				newParams.Add(key, value)
			}
		}

		newParams.Set("page", strconv.Itoa((page - 1)))

		return newParams.Encode()
	}

	return ""
}

func GetNextPage(c echo.Context, page int, endrow int, maxData int64) string {
	if endrow < int(maxData) {
		originalParams := c.QueryParams()

		// Create a new set of query parameters
		newParams := url.Values{}

		// Copy the original query parameters to the new parameters
		for key, values := range originalParams {
			for _, value := range values {
				newParams.Add(key, value)
			}
		}
		newParams.Set("page", strconv.Itoa((page + 1)))

		newParams.Encode()

		return newParams.Encode()
	}

	return ""
}

func GetEndRow(offset int, limit int, totalData int64) int {
	if limit < 0 {
		limit = constant.LIMIT
	}
	if offset+limit <= int(totalData) {
		return limit + offset
	}

	return int(totalData)
}
