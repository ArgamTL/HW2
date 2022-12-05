/*
### Утилита wget
Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	url := flag.String("url", "", "url")
	everything := flag.Bool("o", false, "the whole site")
	flag.Parse()

	urlSlice := strings.Split(*url, "/")
	saveTo := "tempSite/" + urlSlice[len(urlSlice)-1]
	fmt.Println(saveTo)
	resp, err := http.Get(*url)

	if !*everything {

		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		// Create the file
		out, err := os.Create(saveTo)
		if err != nil {
			log.Fatalln(err)
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)

	} else {

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var links []string

		doc.Find("body a").Each(func(index int, item *goquery.Selection) {
			linkTag := item
			link, _ := linkTag.Attr("href")
			links = append(links, link)
		})

		for _, l := range links {
			fmt.Println(l)
			urlSlice := strings.Split(l, "/")
			saveTo := "tempSite/" + urlSlice[len(urlSlice)-1]
			if len(urlSlice) > 2 {
				continue
			}

			resp, err := http.Get(*url + l)
			if err != nil {
				fmt.Println("http.Get *url didn't work")

			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					log.Fatalln(err)
				}
			}()

			f, err := os.Create(saveTo)
			if err != nil {
				fmt.Println("creating file didn't work")
			}
			defer func() {
				err := f.Close()
				if err != nil {
					log.Fatalln(err)
				}
			}()

			_, err = io.Copy(f, resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

		}
	}

}
