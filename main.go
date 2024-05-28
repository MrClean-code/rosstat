package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"regexp"
	"strings"
)

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

	var findLinks func(*html.Node)
	findLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && strings.Contains(attr.Val, "mediabank") {
					matches := re.FindAllString(attr.Val, -1)
					if len(matches) > 0 {
						str := "https://rosstat.gov.ru" + attr.Val
						fmt.Println(str)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findLinks(c)
		}
	}

	findLinks(doc)
}

//
//
//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//
//	"github.com/tebeka/selenium"
//	"github.com/tebeka/selenium/chrome"
//)
//
//type Deal struct {
//	ID           int     `json:"id"`
//	Name         string  `'json:"name"`         // название услуги
//	Owner        string  `'json:"owner"`        // фио или компания
//	Price        int     `'json:"price"`        // цена
//	CountReviews int     `'json:"countReviews"` // количество отзывов
//	Score        float64 `'json:"score"`        // оценка
//	Link         string  `'json:"link"`         // ссылка на продавца
//	// TimeReceivedDeal  string `'json:"timeReceivedDeal"`  // время выполнения
//	// TimeExecutionDeal string `'json:"timeExecutionDeal"` // время получения
//}
//
//var sl2 []Deal
//
//func main() {
//	const chromeDriverPath = "C:\\Program Files\\chromedriver_win32\\chromedriver.exe"
//
//	// Настройка ChromeDriver для использования браузера Chrome
//	fmt.Println("start")
//	opts := []selenium.ServiceOption{}
//	service, err := selenium.NewChromeDriverService(chromeDriverPath, 9515, opts...)
//	if err != nil {
//		fmt.Printf("Ошибка создания сервиса ChromeDriver: %v\n", err)
//		// return
//	}
//	defer service.Stop()
//
//	// Опции браузера Chrome
//	caps := selenium.Capabilities{
//		"browserName": "chrome",
//	}
//	chromeCaps := chrome.Capabilities{
//		Args: []string{
//			// Дополнительные параметры Chrome
//			// Например, можно скрыть окно браузера и запустить в фоновом режиме
//			// "headless",
//			"window-size=1920,1080",
//		},
//	}
//	caps.AddChrome(chromeCaps)
//
//	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
//	if err != nil {
//		fmt.Printf("Ошибка создания WebDriver: %v\n", err)
//		// return
//	}
//	defer wd.Quit()
//
//	url := fmt.Sprintf("https://rosstat.gov.ru/uslugi#")
//	fmt.Println(url)
//
//	fmt.Println("Открытие страницы...")
//	if err := wd.Get(url); err != nil {
//		fmt.Printf("Ошибка открытия страницы: %v\n", err)
//		// return
//	}
//	fmt.Println("Страница успешно открыта")
//
//	rand.Seed(time.Now().UnixNano())
//	randomSeconds := rand.Intn(10) + 9
//	time.Sleep(time.Duration(randomSeconds) * time.Second)
//
//	// elem, err := wd.FindElements(selenium.ByCSSSelector, "a.toggle-card__header-link")
//	// elem2, err := wd.FindElements(selenium.ByCSSSelector, "ul.list-tree__list")
//	elem3, err := wd.FindElements(selenium.ByCSSSelector, "a.btn-sm")
//	if err != nil {
//		fmt.Printf("Ошибка поиска элемента: %v\n", err)
//		return
//	}
//
//	//text, _ := elem[0].Text()
//	//if text == "Платные услуги населению" {
//	//	if err := elem[0].Click(); err != nil {
//	//		fmt.Printf("Ошибка клика по элементу: %v\n", err)
//	//		return
//	//	}
//	//	fmt.Println("Клик по элементу выполнен")
//	//}
//	//
//	//for _, item := range elem {
//	//	text, _ := item.Text()
//	//	if text == "Платные услуги населению" || text == "Бытовые услуги населению" {
//	//		if err := item.Click(); err != nil {
//	//			fmt.Printf("Ошибка клика по элементу: %v\n", err)
//	//			return
//	//		}
//	//		fmt.Println("Клик по элементу выполнен")
//	//	}
//	//}
//	//
//	//for _, el := range elem2 {
//	//	_, err = wd.ExecuteScript("arguments[0].style.display = 'block';", []interface{}{el})
//	//}
//
//	time.Sleep(1 * time.Second)
//
//	fmt.Println("количество ссылок", len(elem3))
//	for _, el := range elem3 {
//		href, _ := el.GetAttribute("href")
//		if href != "" {
//			fmt.Println(href)
//		}
//	}
//	time.Sleep(16 * time.Second)
//	fmt.Println("Успешно")
//}
