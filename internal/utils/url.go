package utils

import "net/url"

func AddFields(fields []string) url.Values {
	q := url.Values{}

	for _, field := range fields {
		q.Add("fields", field)
	}

	return q
}
