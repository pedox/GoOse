package goose

import (
	"github.com/pkg/errors"
)

// Goose is the main entry point of the program
type Goose struct {
	config Configuration
}

// New returns a new instance of the article extractor
func New(args ...string) Goose {
	return Goose{
		config: GetDefaultConfiguration(args...),
	}
}

// ExtractFromURL follows the URL, fetches the HTML page and returns an article object
func (g Goose) ExtractFromURL(url string) (*Article, error) {
	HtmlRequester := NewHtmlRequester(g.config)
	html, err := HtmlRequester.fetchHTML(url)
	if err != nil {
		return nil, errors.Wrap(err, "could not get htnk from site")
	}
	cc := NewCrawler(g.config)
	return cc.Crawl(html, url)
}

// ExtractFromRawHTML returns an article object from the raw HTML content
func (g Goose) ExtractFromRawHTML(RawHTML string, url string) (*Article, error) {
	cc := NewCrawler(g.config)
	return cc.Crawl(RawHTML, url)
}

// ExtractFromURL follows the URL, fetches the HTML page and returns an article object
func (g Goose) ExtractFromURLShort(url string) (*Article, error) {
	HtmlRequester := NewHtmlRequester(g.config)
	html, err := HtmlRequester.fetchHTML(url)
	if err != nil {
		return nil, errors.Wrap(err, "could not get html from from url: "+url)
	}
	cc := NewCrawlerShort(g.config)
	return cc.Crawl(html, url)
}

// ExtractFromRawHTML returns an article object from the raw HTML content
func (g Goose) ExtractFromRawHTMLShort(RawHTML string, url string) (*Article, error) {
	cc := NewCrawlerShort(g.config)
	return cc.Crawl(RawHTML, url)
}
