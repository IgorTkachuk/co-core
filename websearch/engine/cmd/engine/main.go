package main

import (
	"bufio"
	"fmt"
	"go-core/search/crawler/pkg/spider"
	"log"
	"os"
	"strings"
)

type scanner interface {
	Scan(url string, depth int) (data map[string]string, err error)
}

func main() {
	s := new(spider.Instance)

	urls := []string{"https://www.opennet.ru/", "https://go.dev/"}
	res := make(map[string]string)

	res = collect(urls, s)

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter word for search:")
	for input.Scan() {
		for k, v := range res {
			if strings.Contains(v, input.Text()) {
				fmt.Printf("Страница %s имеет адрес: %s\n", v, k)
			}
		}
		fmt.Println("Enter word for search:")
	}

}

func collect(urls []string, sc scanner) (res map[string]string) {
	res = make(map[string]string)
	for _, url := range urls {
		data, err := sc.Scan(url, 2)
		if err != nil {
			log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
		}

		for k, v := range data {
			res[k] = v
		}
	}

	return res
}
