package main

import (
	// "flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gocolly/colly/v2"
	"golang.org/x/exp/slices"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

type PageData struct {
	FeedLinks []string
}

func main() {
	var url, sp string
	flag.StringVar(&url, "l", "http://jsonfeed.org", "Enter link for parsing rss feeds")
	flag.StringVar(&sp, "s", "fast", "Enter a value for setting speed of scraping feeds")
	flag.Parse()
	if url != "" {
		ParseData(url, sp)
	} else {
		http.HandleFunc("/", handler)
		http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
		log.Fatal(http.ListenAndServe(":5000", nil))
	}
	
}

func handler (w http.ResponseWriter, r *http.Request) {
	var ch = make(chan []string, 100)
	u, err := url.Parse(r.URL.String())
	var data  = PageData{}
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    } else {
		params := u.Query()
		url := params.Get("url")
		sp := params.Get("sp")
		// result := params.Get("result")
		if url != "" {
			url = FormatHostUrl(url)
			go getData(url,  sp, ch)
			// linksSet := mapset.NewSet[string]()
			links := <-ch
			fmt.Println(links)
			linksSet := mapset.NewSet[string]()
			for _, link := range links {
				if link != "broken_link" {
					linksSet.Add(link)
				}
			}
			links = linksSet.ToSlice()
			slices.Sort(links)

			data = PageData {
					FeedLinks: links,
			}
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func ParseData(url, sp string, ch chan []string) {
	c := colly.NewCollector(colly.Async())	
	
	var links []string

	c.OnHTML("head", func(h *colly.HTMLElement) {
		if len(links) == 0 {
			h.ForEach("link", func(i int, ch *colly.HTMLElement) {
				link := ch
				attr := link.Attr("type")
				if m, _ := regexp.MatchString("atom|rss", attr); m {
					href := link.Attr("href")
					fmt.Println(href)
					link := FormatLink(url, href)
					fmt.Println(link)
					links = append(links, link)
				}
			})
		}
	})

	if sp != "fast" {
		c.OnHTML("a", func(h *colly.HTMLElement) {
			attr := h.Attr("href")
			link := FormatLink(url, attr)
			if !strings.Contains(attr, "cgi") {
				if m, _ := regexp.MatchString(`\.rss$|\/feed$`, attr); m {
					fmt.Println(link)
					links = append(links, link)
				} else {
					if len(links) == 0 {
						h.Request.Visit(link)
					}
				}
			}
		})
	}

	fmt.Println(links)

	c.Visit(url)

	c.Wait()

	if len(links) > 0 {
		ch <- links
	}
}

func FormatHostUrl(url string) string {
	if !strings.HasPrefix(url, "http") {
		url = `http://` + url
	}
	return url
}

func FormatLink(url, link string) string {
	var ref string
	if strings.HasPrefix(link, "http") {
		ref = link
	}else if strings.HasPrefix(link, "//") {
		ref = fmt.Sprintf("http:%s", link)
	} else if m, _ := regexp.MatchString(`^(\.)|(\/{1}\w+)`, link); m {
		if strings.HasPrefix(link, ".") {
			ref = url + link[1:]
		} else {
			ref = url + link
		}
	}
	if ref != "" {
		return ref
	}
	return "broken_link"
}