package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type problem struct {
	content string
	id      string
}

var myLogger *log.Logger

func basicRouter(w http.ResponseWriter, r *http.Request) {
	var localPath string = "public/"
	myLogger.Print(r.URL)
	content, err := ioutil.ReadFile(localPath + r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}
	w.Write(content)
}

func crowrlingPage(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}

func main() {
	myLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	http.HandleFunc("/", basicRouter)
	if error := http.ListenAndServe(":3000", nil); error != nil {
		panic(error)
	}
}

/*
	nil 은 포인터, 인터페이스, 맵, 슬라이스, 채널, 함수의 zero value 입니다.


*/
