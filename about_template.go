package koans

import (
	"bytes"
	"html/template"
	"strings"
)

func aboutTemplate() {
	{
		// Give primitive value
		templateStr := `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8" />
<title>{{ . }}</title>
</head>
<body>
</body>
</html>
`
		tmpl := template.Must(template.New("qr").Parse(templateStr))
		b := new(bytes.Buffer)
		tmpl.Execute(b, "Hello World!")
		assert(strings.Contains(b.String(), stringstring))
	}

	{
		// Give Struct value
		templateStr := `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8" />
<title>Hello World</title>
</head>
<body>
    <ul>
        {{ range .Scores }}
        <li>{{ . }}</li>
        {{ end }}
    </ul>
    <h1>{{ .Name }}</h1>
    <h2>{{ .age }}</h2>
</body>
</html>
`
		student := struct {
			Scores []int
			Name   string
			age    int
		}{
			Scores: []int{80, 30, 50},
			Name:   "Hoge",
			age:    30,
		}
		tmpl := template.Must(template.New("qr").Parse(templateStr))
		b := new(bytes.Buffer)
		tmpl.Execute(b, student)
		assert(strings.Count(b.String(), "<li>") == intintint)
		assert(strings.Contains(b.String(), stringstring))
		assert(strings.Contains(b.String(), stringstring))
	}
}
