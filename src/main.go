package main

import (
	"io/ioutil"
	"net/http"
)

type problem struct {
	content string
	id      string
}

func basicRouter(w http.ResponseWriter, r *http.Request) {
	var localPath string = "public/"
	content, err := ioutil.ReadFile(localPath + r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}
	w.Write(content)
}

func main() {
	http.HandleFunc("/", basicRouter)
	if error := http.ListenAndServe(":3000", nil); error != nil {
		panic(error)
	}
}

/*
	nil 은 포인터, 인터페이스, 맵, 슬라이스, 채널, 함수의 zero value 입니다.


*/
