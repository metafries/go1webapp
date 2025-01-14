package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// url := "http://localhost:8000/1ei30n11"
	// resp, err := http.Get(url)
	// inspectResponse(resp, err)

	// data, err := json.Marshal(struct {
	// 	X int
	// 	Y float32
	// }{X: 4, Y: 3.8})
	// if err != nil {
	// 	log.Fatal("FATAL: Error Occured While Marshaling JSON ", err)
	// }
	// resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	// inspectResponse(resp, err)

	// client := http.Client{
	// 	Timeout: 3 * time.Second,
	// }
	// client.Get(url)

	// req, err := http.NewRequest(http.MethodPut, url, nil)
	// if err != nil {
	// 	log.Fatal("FATAL: ", err)
	// }
	// req.Header.Add("x-testheader", "Testing Add Header")
	// req.Header.Set("User-Agent", "Testing Set Header")
	// resp, err := client.Do(req)
	// inspectResponse(resp, err)

	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		log.Fatal("FATAL: ", err)
	}
	defer resp.Body.Close()
	v := struct {
		IP string `json:"ip"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		log.Fatal("FATAL: ", err)
	}
	log.Println(v.IP)
}

func inspectResponse(resp *http.Response, err error) {
	if err != nil {
		log.Fatal("FATAL: Error Occured While Marshaling JSON ", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("FATAL: Error Occured While Trying to Read HTTP Response Body ", err)
	}
	log.Println(string(b))
}
