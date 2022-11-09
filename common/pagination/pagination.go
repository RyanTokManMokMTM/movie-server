package pagination

import "math"

const (
	MAX_LIMIT         = 20
	MAX_OFFSET        = 20
	DEFAULT_PAGE_SIZE = 20
)

func GetOffset(offset uint) uint {
	if offset > MAX_OFFSET {
		return MAX_OFFSET
	} else if offset < 0 {
		return 0
	}

	return offset
}

func GetLimit(limit uint) uint {
	if limit > MAX_LIMIT {
		return MAX_LIMIT
	} else if limit < 0 {
		return 0
	}

	return limit
}

func GetPage(page uint) uint {
	if page < 0 {
		return 0
	}
	return page
}

func PageOffset(pageSize, page uint) uint {
	return (pageSize) * (page - 1)
}

func GetTotalPageByPageSize(total, pageSize uint) uint {
	return uint(math.Ceil(float64(total) / float64(pageSize)))
}
