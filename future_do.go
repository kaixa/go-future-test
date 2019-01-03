// maingo
package main

import (
	"fmt"
	"io/ioutil"
	"time"
	//"log"
	"net/http"
)

func main() {
	useUrl := "https://www.baidu.com"
	fmt.Println("Beging...")
	fmt.Println("request:", useUrl)
	future := RequestFuture(func() (interface{}, error) {
		c := &http.Client{
			Timeout: 500 * time.Millisecond,
		}
		resp, err := c.Get(useUrl)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	})

	backResult, backError := future()
	body, _ := backResult.([]byte)
	fmt.Println("reponse length:", len(body))
	fmt.Println("reponse length:", backError)
}

/**
 * RequestFuture, http request promise.
 *
 */
func RequestFuture(f func() (interface{}, error)) func() (interface{}, error) {
	var result interface{}
	var err error

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)
		result, err = f()
	}()

	return func() (interface{}, error) {
		<-c
		return result, err
	}
}
