package client

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	fetchCoverageURL = "https://api.openbd.jp/v1/coverage"
	fetchBookInfoURL = "https://api.openbd.jp/v1/get"
)

// FetchCoverage OpenBDに登録されている書誌のISBNを取得
func FetchCoverage() (string, error) {
	resp, err := http.Get(fetchCoverageURL)

	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FetchBookInfo OpenDB APIからISBNをもとに書誌データを取得
func FetchBookInfo(i []string) (string, error) {
	v := url.Values{}
	v.Add("isbn", strings.Join(i, ","))

	resp, err := retry(func() (*http.Response, error) {
		resp, err := http.PostForm(fetchBookInfoURL, v)
		return resp, err
	})

	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func retry(f func() (*http.Response, error)) (*http.Response, error) {
	var err error
	for i := 0; i < 3; i++ {
		resp, err := f()
		if err == nil {
			return resp, nil
		}
	}
	return nil, err
}
