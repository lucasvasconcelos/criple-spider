package parse

import (
	"crypto/tls"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"time"
)

// HTMLPage type to be used by the parser
type HTMLPage struct {
	Url      string
	Links    []string
	response *http.Response
}

// Create a new HTMLPage
func NewHTMLPage(url string) *HTMLPage {
	h := &HTMLPage{
		Url: url,
	}
	return h
}

// Populates HTMLPage.response
func (h *HTMLPage) getResponse() error {
	if h.response != nil {
		return nil
	}

	httpClient := &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	response, err := httpClient.Get(h.Url)

	if err != nil {
		return err
	}

	h.response = response

	return nil
}

// Parse the links inside a HTMLPage body
func (h *HTMLPage) ParseLinks() error {
	err := h.getResponse()

	if err != nil {
		return err
	}

	defer h.response.Body.Close()

	tokens := html.NewTokenizer(h.response.Body)

	for {
		token := tokens.Next()

		// TODO: treat token erros
		if token == html.ErrorToken {
			break
		}

		tagName, hasAttr := tokens.TagName()

		if len(tagName) == 1 && tagName[0] == 'a' && hasAttr {
			key, value, _ := tokens.TagAttr()
			if string(key) == "href" && strings.HasPrefix(string(value), "http") {
				h.Links = append(h.Links, string(value))
			}
		}
	}

	return nil
}
