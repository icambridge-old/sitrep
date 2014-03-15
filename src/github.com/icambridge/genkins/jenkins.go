package genkins

import (
  "log"
	"io/ioutil"
  "net/http"
	"encoding/json"
)

func GetJobView() (jobView *JobView, err error)  {
  url := "http://localhost:8080/" + "api/json?tree=jobs[name,url,color]"

	
	res, err := http.Get(url)

	if err != nil {
		log.Println("Unable to connect to Jenkins")
		return nil, err
	}
	bs, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println("Invalid HTTP response from Jenkins?!")
		return nil, err
	}

	tr := string(bs)
  var view JobView
	err = json.Unmarshal([]byte(tr), &view)

	if err != nil {

		log.Println("Can't unmarshal JSON response from Jenkins")
		return nil, err
	}

	return &view, nil
}
