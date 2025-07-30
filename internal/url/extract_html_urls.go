package url

import (

	"strings"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {

	ans := make([]string,0)
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {

		return nil, err
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					
					if a.Val[0] == '/' {

						ans = append(ans, rawBaseURL + a.Val)
					} else {

						ans = append(ans, a.Val);
					}
					break
				}
			}
		}
	}

	return ans, nil
	
}
