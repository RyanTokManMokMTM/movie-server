package pagination

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOffset(t *testing.T) {
	testCases := []struct {
		Name     string
		offset   uint
		expected uint
	}{
		{
			Name:     "Normal offset #1",
			offset:   10,
			expected: 10,
		},
		{
			Name:     "Normal offset #1",
			offset:   12,
			expected: 12,
		},
		{
			Name:     "Normal offset #1",
			offset:   3,
			expected: 3,
		},
		{
			Name:     "over max offset #1",
			offset:   25,
			expected: 20,
		},
		{
			Name:     "over max offset #2",
			offset:   50,
			expected: 20,
		},
		{
			Name:     "over max offset #3",
			offset:   10000,
			expected: 20,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			offset := GetOffset(test.offset)
			assert.Equal(t, test.expected, offset)
		})
	}
}

func TestGetLimit(t *testing.T) {
	testCases := []struct {
		Name     string
		limit    uint
		expected uint
	}{
		{
			Name:     "Normal offset #1",
			limit:    10,
			expected: 10,
		},
		{
			Name:     "Normal offset #1",
			limit:    12,
			expected: 12,
		},
		{
			Name:     "Normal offset #1",
			limit:    3,
			expected: 3,
		},
		{
			Name:     "over max offset #1",
			limit:    25,
			expected: 20,
		},
		{
			Name:     "over max offset #2",
			limit:    50,
			expected: 20,
		},
		{
			Name:     "over max offset #3",
			limit:    10000,
			expected: 20,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			offset := GetLimit(test.limit)
			assert.Equal(t, test.expected, offset)
		})
	}
}

func TestGetPage(t *testing.T) {
	testCases := []struct {
		Name     string
		page     uint
		expected uint
	}{
		{
			Name:     "Test Get Page #1",
			page:     1,
			expected: 1,
		},
		{
			Name:     "Test Get Page #2",
			page:     10,
			expected: 10,
		},
		{
			Name:     "Test Get Page #3",
			page:     100,
			expected: 100,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			p := GetPage(test.page)
			assert.Equal(t, test.expected, p)
		})
	}
}

func TestPageOffset(t *testing.T) {
	testCases := []struct {
		Name     string
		PageSize uint
		page     uint
		expected uint
	}{
		{
			Name:     "Test Get Page Offset #1",
			PageSize: 20,
			page:     1,
			expected: 0,
		},
		{
			Name:     "Test Get Page Offset #2",
			PageSize: 20,
			page:     2,
			expected: 20,
		},
		{
			Name:     "Test Get Page Offset #3",
			PageSize: 20,
			page:     10,
			expected: 180,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			pageOffset := PageOffset(test.PageSize, test.page)
			assert.Equal(t, test.expected, pageOffset)
		})
	}
}

func TestGetTotalPageByPageSize(t *testing.T) {
	testCases := []struct {
		Name     string
		Total    uint
		PageSize uint
		expected uint
	}{
		{
			Name:     "Test Get Total Page #1",
			Total:    100,
			PageSize: 20,
			expected: 5,
		},
		{
			Name:     "Test Get Total Page #2",
			Total:    100,
			PageSize: 10,
			expected: 10,
		},
		{
			Name:     "Test Get Total Page #3",
			Total:    333,
			PageSize: 20,
			expected: 17,
		},
		{
			Name:     "Test Get Total Page #3",
			Total:    6951,
			PageSize: 10,
			expected: 696,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			page := GetTotalPageByPageSize(test.Total, test.PageSize)
			assert.Equal(t, test.expected, page)
		})
	}
}
