package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/luabagg/orcgen/v2"
	"github.com/luabagg/orcgen/v2/pkg/handlers/pdf"
)

var server = http.Server{
	Addr: ":8080",
}

func main() {

	var file1 = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>File 1</title>
</head>
<body>
    <h1>File 1 test</h1>
</body>
</html>
`
	var file2 = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>File 2</title>
</head>
<body>
    <h1>File 2 test</h1>
</body>
</html>
`

	var file3 = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>File 3</title>
</head>
<body>
    <h1>File 3 test</h1>
</body>
</html>
`

	files := [3]string{file1, file2, file3}

	for n, f := range files {
		fNameHTML := fmt.Sprintf("file%d.html", n+1)
		fNamePDF := fmt.Sprintf("file%d.pdf", n+1)
		os.WriteFile(fNameHTML, []byte(f), 0644)

		htmlFile, err := os.Open(fNameHTML)
		if err != nil {
			log.Fatalf("opening HTML file: %w", err)
		}
		defer htmlFile.Close()

		html, err := io.ReadAll(htmlFile)
		if err != nil {
			log.Fatalf("could not read the HTML file: %w", err)
		}

		// Setup web server in the background
		// for the PDF rendering functionality to act as a headless browser
		serve(html)

		pdf, err := orcgen.ConvertWebpage(pdf.New().SetConfig(orcgen.PDFConfig{
			DisplayHeaderFooter: false,
			PreferCSSPageSize:   true,
		}), "http://localhost:8080")
		if err != nil {
			log.Fatalf("converting HTML to PDF: %v", err)
		}
		pdf.Output(fNamePDF)

		ctx := context.WithoutCancel(context.Background())
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}

	}
}

func serve(content []byte) {
	mux := http.NewServeMux()
	mux.Handle("/", contentHandler(content))
	server.Handler = mux

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server started.")
}

func contentHandler(content []byte) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Write the HTML content to the response
		w.Header().Set("Content-Type", "text/html")
		w.Write(content)
	}
	return http.HandlerFunc(fn)
}
