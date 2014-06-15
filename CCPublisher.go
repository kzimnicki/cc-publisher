package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"log"
	"net/http"
	"net/url"
)

func translate(b byte) string {

	germanLetters := map[byte]string{
		0x7b: "ä",
		0x7d: "ü",
		0x7c: "ö",
		0x7e: "ß",
	}

	b &= 0x7F
	var letter = ""
	if val, ok := germanLetters[b]; ok {
		letter = val
	} else {
		letter = string(b)
	}
	//	fmt.Printf(letter)
	return letter
}

func httpPost(text string) {
	log.Println(text)
	resp, err := http.PostForm("http://explain.cc:8888/app/", url.Values{"v": {text}})
	if err != nil {
//		log.Fatalln(err)
	} else {
		resp.Close = true
		defer resp.Body.Close()
	}
}

func process() {

	r := bufio.NewReader(os.Stdin)

	var text = false
	var counter = 0
	var printed = false
	var subtitle = ""
	var doubleLine = false

	for true {
		data := make([]byte, 200)
		data = data[:cap(data)]
		n, err := r.Read(data)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				return
			}
			fmt.Println(err)
		}
		data = data[:n]
		for _, b := range data {
			if (b == 0x02 || b == 0x0d || b == 0x83) {
				continue;
			}
			if b == 0x8a {
				counter++
				if (!printed && !doubleLine) {
					//					fmt.Println("|"+strings.TrimSpace(subtitle)+"|")
					httpPost(strings.TrimSpace(subtitle))
					printed = true
					subtitle = ""
				}
				text = false
			}

			if text {
				subtitle += translate(b)
			}

			if b == 0x9b || b == 0x8c {
				doubleLine = (b == 0x8c)
				//				fmt.Printf("Start")
				text = true
				printed = false;
			}
		}
	}
}

func main() {
	fmt.Println("START")
	process()
}
