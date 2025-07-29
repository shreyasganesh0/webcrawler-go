package test


import ()


func TestGetURLSFromHTML(t *testing.T) {

	test_cases := struct{

		name      string
		inputURL  string
		inputBody string
		expected  []string

	}{
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

}
