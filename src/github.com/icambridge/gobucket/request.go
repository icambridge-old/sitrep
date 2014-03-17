package gobucket

import (
	"io"
	"net/http"
	"log"
	"strings"
)

type Client struct {
  request *Request
   
}

func NewRequest(uname string, pword string) *Request {
  return &Request{Username: uname, Password: pword}
}


type Request struct {
	Username string
	Password string
	
}

func (r Request) get(url string) *http.Response {
	req, err := http.NewRequest("GET", url, strings.NewReader(""))

	if err != nil {
		log.Fatal(err)
	}

	return r.process(req)
}

func (r Request) postWithBody(url string, body io.Reader, contentType string) *http.Response {
	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", contentType)

	return r.process(req)
}

func (r Request) post(url string) *http.Response {
	req, err := http.NewRequest("POST", url, strings.NewReader(""))

	if err != nil {
		log.Fatal(err)
	}

	return r.process(req)
}

func (r Request) put(url string) *http.Response {
	req, err := http.NewRequest("PUT", url, strings.NewReader(""))

	if err != nil {
		log.Fatal(err)
	}

	return r.process(req)
}

func (r Request) delete(url string) *http.Response {
	req, err := http.NewRequest("DELETE", url, strings.NewReader(""))

	if err != nil {
		log.Fatal(err)
	}

	return r.process(req)
}

func (r Request) process(req *http.Request) *http.Response {

  req.SetBasicAuth(r.Username, r.Password)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	return res
}
