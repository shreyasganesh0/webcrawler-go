package test


import (
	"testing"
	"reflect"
	"github.com/shreyasganesh0/webcrawler-go/internal/url"
)


func TestGetURLSFromHTML(t *testing.T) {

	test_cases := []struct{

		name      string
		inputURL  string
		inputBody string
		expected  []string

	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
				expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
	}

	for i, tc := range test_cases {

		t.Run(tc.name, func(t *testing.T) {
			actual, err := url.GetURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}

}
