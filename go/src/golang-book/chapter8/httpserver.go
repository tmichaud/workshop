package main

import (
	"net/http"
	"io"
)


// Obviously, test this with a web browser by going to http://localhost:9000/hello
func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type", 
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
		  <head>
		    <title>Hello World</title>
		  </head>
		  <body>
		    Hello World...from Luna.
		  </body>
		</html>`, 
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
	http.Handle (
		"/assets/", 
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("./assets/")),
		),
	)
}
