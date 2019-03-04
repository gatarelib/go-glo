package utils

import "net/url"

// AddFields used to help construct query fields
// for API calls.
func AddFields(fields []string) url.Values {
	q := url.Values{}

	for _, field := range fields {
		q.Add("fields", field)
	}

	return q
}
