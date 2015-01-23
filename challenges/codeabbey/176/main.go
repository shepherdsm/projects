package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Say 100
// http://www.codeabbey.com/index/task_view/say-100
func main() {
	url := "http://codeabbey.sourceforge.net/say-100.php"
	content := "text/plain"
	body := "token: Bz8ZOkK2Of70emxLW8RYIjUz"

	resp := makeReq(url, content, strings.NewReader(body))

	var secret int
	fmt.Sscanf(resp, "secret: %d", &secret)

	body = fmt.Sprintf("%s\nanswer: %d", body, 100-secret)

	resp = makeReq(url, content, strings.NewReader(body))

	fmt.Println(resp)
}

func makeReq(url, content string, body io.Reader) string {
	resp, _ := http.Post(url, content, body)
	defer resp.Body.Close()

	read, _ := ioutil.ReadAll(resp.Body)
	return string(read)
}
