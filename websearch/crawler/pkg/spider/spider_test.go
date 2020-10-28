// Package spider реализует сканер содержимого веб-сайтов.

// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.

package spider

import (
	"testing"
)

type scanner interface {
	Scan(url string, depth int) (data map[string]string, err error)
}

func TestScanSite(t *testing.T) {
	s := new(TScan)
	url := "https://www.opennet.ru/"
	data, err := collect(url, s)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range data {
		t.Logf("%s -> %s\n", k, v)
	}
}

func TestStubScanSite(t *testing.T) {
	s := new(Testscan)

	data, err := collect("", s)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range data {
		t.Logf("%s -> %s\n", k, v)
	}
}

func collect(url string, sc scanner) (res map[string]string, err error) {
	res = make(map[string]string)

	data, err := sc.Scan(url, 2)
	if err != nil {
		return data, err
	}

	return data, nil
}
