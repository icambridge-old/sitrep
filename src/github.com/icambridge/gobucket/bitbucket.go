package gobucket

type Client struct {
	Request     *Request
	
	Repository  *RepositoryService
	
}

func GetClient(username string, password string) *Client {
  r := &Request{Username: username, Password: password}
  c := &Client{Request: r}
  c.Repository = &RepositoryService{Client: c}
  return c
}
