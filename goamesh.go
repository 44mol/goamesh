package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main () {
	showText("End!", 5 * time.Second)
}

func showText (text string, limit time.Duration) {
	clearScreen()

	for begin := time.Now(); time.Since(begin) < limit; {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(1 * time.Second)
		clearScreen()
	}
	fmt.Println(text)
}

func clearScreen () {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getPic () {
	var url string = "https://tokyo-ame.jwa.or.jp/mesh/050/201907281810.gif"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("status:", response.Status)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("hoge", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer func() {
		file.Close()
	}()

	file.Write(body)
}
