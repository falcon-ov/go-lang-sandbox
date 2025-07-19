package serversproblems

import (
	"io"
	"net/http"
)

func Hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>

<html>

	<head>
	    <title>Hello World</title>
	</head>
	<body>
	    Hello World GOLANG!
	</body>

</html>`,
	)
}
