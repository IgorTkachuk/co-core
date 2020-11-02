package spider

// StubInstance - заглушка сканера с реализованным методом Scan
type StubInstance struct{}

// Scan - метод возвращает заранее предопределенный результат сканирования
func (s *StubInstance) Scan(url string, depth int) (data map[string]string, err error) {
	data = make(map[string]string)

	data["addr1"] = "name addr 1"
	data["addr2"] = "name addr 2"
	data["addr3"] = "name addr 3"
	data["addr4"] = "name addr 4"
	data["addr5"] = "name addr 5"

	return data, nil
}
