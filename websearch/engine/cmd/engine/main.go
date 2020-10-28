package main

import (
	"flag"
	"fmt"
	"go-core/search/crawler/pkg/spider"
	"log"
	"strings"
)

func main() {
	urls := []string{"https://www.opennet.ru/", "https://go.dev/"}
	res := make(map[string]string)

	for _, url := range urls {
		data, err := spider.Scan(url, 2)
		if err != nil {
			log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
		}

		for k, v := range data {
			res[k] = v
		}
	}

	s := flag.String("s", "", "word for search")
	flag.Parse()

	if *s == "" {
		fmt.Println("Nothing to search")
		return
	}

	for k, v := range res {
		if strings.Contains(v, *s) {
			fmt.Printf("Страница %s имеет адрес: %s\n", v, k)
		}
	}
}
