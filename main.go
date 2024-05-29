package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func downloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при загрузке файла: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("не удалось скачать файл: статус-код %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении содержимого файла: %v", err)
	}

	return body, nil
}

func main() {
	url := "https://rosstat.gov.ru/uslugi"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при получении страницы:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Статус-код не является 200 OK")
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при парсинге HTML:", err)
		return
	}

	re := regexp.MustCompile(`\d{6}`)
	var links []string
	var findLinks func(*html.Node)
	findLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && strings.Contains(attr.Val, "mediabank") {
					matches := re.FindAllString(attr.Val, -1)
					if len(matches) > 0 {
						str := "https://rosstat.gov.ru" + attr.Val
						links = append(links, str)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findLinks(c)
		}
	}

	findLinks(doc)

	for _, link := range links {
		fileContent, err := downloadFile(link)
		if err != nil {
			fmt.Printf("Ошибка при скачивании файла %s: %v\n", link, err)
			continue
		}
		fmt.Printf("Скачан файл: %s (размер: %d байт)\n", link, len(fileContent))

		// Сохранение файла на диск
		filename := link[strings.LastIndex(link, "/")+1:]
		err = ioutil.WriteFile(filename, fileContent, 0644)
		if err != nil {
			fmt.Printf("Ошибка при сохранении файла %s: %v\n", filename, err)
		}
	}
}
