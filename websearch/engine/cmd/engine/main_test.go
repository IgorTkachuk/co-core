package main

import (
	"go-core/search/crawler/pkg/spider"
	"testing"
)

func TestCollect(t *testing.T) {
	//Используем заглушку из пакетв spider. Не используем сетевое взаимодействие
	s := new(spider.StubInstance)
	urls := []string{""}

	data := collect(urls, s)

	want := "name addr 3"
	got := data["addr3"]

	if want != got {
		t.Fatal("Ошибка тестировании функции collect")
	}
}
