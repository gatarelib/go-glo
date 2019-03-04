package glo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (a *Glo) multiPartReq(
	method string,
	url string,
	r io.Reader,
	q url.Values) (
	data []byte,
	header http.Header,
	err error,
) {
	return a.do(method, url, r, q, "multipart/form-data")
}

func (a *Glo) jsonReq(
	method string,
	url string,
	b []byte,
	q url.Values,
) (
	data []byte,
	header http.Header,
	err error,
) {
	return a.do(method, url, bytes.NewReader(b), q, "application/json")
}

func (a *Glo) do(
	method string,
	url string,
	r io.Reader,
	q url.Values,
	contentType string,
) (
	data []byte,
	header http.Header,
	err error,
) {
	req, err := http.NewRequest(method, url, r)
	if err != nil {
		err = fmt.Errorf("failed to construct request err:%s", err)
		return
	}
	err = setRequestHeaders(req, a.token, contentType)
	if err != nil {
		return
	}
	req.URL.RawQuery = q.Encode()

	resp, err := a.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	header = resp.Header

	switch resp.StatusCode {
	case http.StatusOK:
		data, err = ioutil.ReadAll(resp.Body)
	case http.StatusNoContent:
	case http.StatusTooManyRequests:
		err = fmt.Errorf(
			"rate limit reached",
		)
		return
	default:
		err = fmt.Errorf(
			"unsupported response httpCode:%d status:%s ",
			resp.StatusCode,
			resp.Status,
		)
		return
	}

	return
}

func setRequestHeaders(
	req *http.Request,
	token string,
	contentType string,
) (
	err error,
) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", contentType)

	return
}
