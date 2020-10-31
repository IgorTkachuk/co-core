// Package spider реализует сканер содержимого веб-сайтов.

// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.

package spider

import (
	"testing"
)

func TestScan(t *testing.T) {
	const url = "https://www.opennet.ru/"
	const depth = 2
	s := new(TestInstance)

	data, err := s.Scan(url, depth)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range data {
		t.Logf("%s -> %s\n", k, v)
	}
}
