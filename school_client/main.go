package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	testUrl := "http://www.baidu.com"
	client := http.Client{Timeout: time.Second * 10}
	request, e := http.NewRequest(http.MethodGet, testUrl, nil)
	handlerError(e)
	response, e := client.Do(request)
	handlerError(e)
	readCloser := response.Body
	defer func() {
		handlerError(readCloser.Close())
	}()
	bodyBytes, e := ioutil.ReadAll(readCloser)
	handlerError(e)
	body := string(bodyBytes)
	firstIndex := strings.Index(body, "='")
	lastIndex := strings.LastIndex(body, "'<")
	location := body[firstIndex+2 : lastIndex]
	if strings.HasPrefix(location, "http") {
		fmt.Println(location)
		wg.Add(1)
		go openBrowser(location)
		wg.Wait()
	} else {
		fmt.Println("no need login")
	}
}

func handlerError(e error) {
	if e != nil {
		log.Printf("Have Error %s", e.Error())
		panic(e)
	}
}

func openBrowser(url string) {
	fmt.Println("请稍后...")
	fmt.Println("打开浏览器...")
	openError := open(url)
	if openError != nil {
		fmt.Println(openError)
	}
	wg.Done()
}

// Open calls the OS default program for uri
func open(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}
