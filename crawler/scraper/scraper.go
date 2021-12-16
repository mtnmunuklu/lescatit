package scraper

import (
	"Lescatit/pb"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type CollyScraper interface {
	GetLinks(url string) ([]string, error)
	GetData(url string) (string, error)
	FromProtoBuffer(crawler *pb.CrawlRequest)
}

type CScraper struct {
	userAgent            string
	maxDepth             int64
	maxBodySize          int64
	allowedDomains       []string
	disallowedDomains    []string
	disallowedURLFilters []*regexp.Regexp
	urlFilters           []*regexp.Regexp
	urlRevisit           bool
	robotsTxt            bool
}

// NewCollyScraper creates a new CollyScraper instance with default configuration.
func NewCollyScraper() CollyScraper {
	cs := &CScraper{}
	cs.Init()
	return cs
}

// Init initializes the CScraper's private variables and sets default
// configuration for the CSraper
func (cs *CScraper) Init() {
	cs.userAgent = "colly - https://github.com/gocolly/colly"
	cs.maxDepth = 0
	cs.maxBodySize = 10 * 1024 * 1024
	cs.urlRevisit = false
	cs.robotsTxt = true
}

//GetLinks provides to get the links in the url address.
func (cs *CScraper) GetLinks(url string) ([]string, error) {
	var links []string
	collector := colly.NewCollector(
		colly.UserAgent(cs.userAgent),
		colly.MaxDepth(int(cs.maxDepth)),
		colly.MaxBodySize(int(cs.maxBodySize)),
		colly.AllowedDomains(cs.allowedDomains...),
		colly.DisallowedDomains(cs.disallowedDomains...),
		colly.DisallowedURLFilters(cs.disallowedURLFilters...),
		colly.URLFilters(cs.urlFilters...),
		func(c *colly.Collector) {
			c.AllowURLRevisit = cs.urlRevisit
		},
		func(c *colly.Collector) {
			c.IgnoreRobotsTxt = cs.urlRevisit
		},
	)
	collector.SetRequestTimeout(3 * time.Second)

	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		url := e.Request.AbsoluteURL(e.Attr("href"))
		if url != "" {
			links = append(links, strings.TrimSpace(url))
		}
	})

	collector.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL)
	})

	err := collector.Visit(url)
	if err != nil {
		return nil, err
	}

	return links, nil
}

//GetData provides to get the content in the url address.
func (cs *CScraper) GetData(url string) (string, error) {
	var data []string
	collector := colly.NewCollector()
	collector.SetRequestTimeout(3 * time.Second)

	collector.OnHTML("h", func(e *colly.HTMLElement) {
		content := e.Text
		if content != "" {
			data = append(data, content)
		}
	})
	collector.OnHTML("p", func(e *colly.HTMLElement) {
		content := e.Text
		if content != "" {
			data = append(data, content)
		}
	})

	err := collector.Visit(url)
	if err != nil {
		return "", err
	}

	return strings.Join(data, " "), nil
}

// FromProtoBuffer gets data from protocol buffer and converts to the cscraper structure.
func (cs *CScraper) FromProtoBuffer(crawler *pb.CrawlRequest) {
	if crawler.GetUserAgent() != "" {
		cs.userAgent = crawler.GetUserAgent()
	}

	cs.maxDepth = crawler.GetMaxDepth()
	if crawler.GetMaxBodySize() != 0 {
		cs.maxBodySize = crawler.GetMaxBodySize()
	}

	cs.allowedDomains = crawler.GetAllowedDomains()
	cs.disallowedDomains = crawler.GetDisallowedDomains()
	for _, duf := range crawler.GetDisallowedUrlFilters() {
		cs.disallowedURLFilters = append(cs.disallowedURLFilters, regexp.MustCompile(duf))
	}

	for _, uf := range crawler.GetUrlFilters() {
		cs.urlFilters = append(cs.urlFilters, regexp.MustCompile(uf))
	}

	cs.urlRevisit = crawler.GetUrlRevisit()
	cs.robotsTxt = crawler.GetRobotsTxt()
}
