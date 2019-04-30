package apiutil

import (
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestPageFromRequest(t *testing.T) {
	cases := []struct {
		description        string
		queries            map[string]int
		expNumber, expSize int
	}{
		{"no page or pageSize params", map[string]int{}, 1, DEFAULT_PAGE_SIZE},
		{"negative ints for page and pageSize params", map[string]int{"page": -1, "pageSize": -1}, 1, DEFAULT_PAGE_SIZE},
		{"no pageSize param", map[string]int{"page": 2}, 2, DEFAULT_PAGE_SIZE},
		{"no page param", map[string]int{"pageSize": 25}, 1, 25},
		{"happy path", map[string]int{"page": 5, "pageSize": 30}, 5, 30},
	}

	for _, c := range cases {
		r := httptest.NewRequest("GET", "/", nil)
		q := r.URL.Query()
		// add query params
		for key, val := range c.queries {
			q.Set(key, strconv.Itoa(val))
		}
		r.URL.RawQuery = q.Encode()

		got := PageFromRequest(r)
		if c.expNumber != got.Number {
			t.Errorf("case '%s' error: number mismatch, expected '%d', got '%d'", c.description, c.expNumber, got.Number)
		}
		if c.expSize != got.Size {
			t.Errorf("case '%s' error: size mismatch, expected '%d', got '%d'", c.description, c.expSize, got.Size)
		}
	}

}

func TestNewPageFromLimitAndOffset(t *testing.T) {
	cases := []struct {
		description                       string
		offset, limit, expNumber, expSize int
	}{
		{"offset and limit 0", 0, 0, 1, DEFAULT_PAGE_SIZE},
		{"offset and limit negative", -1, -1, 1, DEFAULT_PAGE_SIZE},
		{"offset and limit happy path", 150, 25, 7, 25},
		{"offset and limit offset not multiple of limit", 90, 25, 4, 25},
		{"offset and limit larger limit then offset", 25, 150, 1, 150},
	}

	for _, c := range cases {
		got := NewPageFromOffsetAndLimit(c.offset, c.limit)
		if c.expNumber != got.Number {
			t.Errorf("case '%s' error: number mismatch, expected '%d', got '%d'", c.description, c.expNumber, got.Number)
		}
		if c.expSize != got.Size {
			t.Errorf("case '%s' error: size mismatch, expected '%d', got '%d'", c.description, c.expSize, got.Size)
		}
	}
}
