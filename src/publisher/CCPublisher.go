package publisher

import (
	"bufio"
	"fmt"
	"time"
	"io"
	"os"
	"strings"
//	"log"
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

func httpPost(line1 string, line2 string) {
	fmt.Println(len(line1), line1)
//	log.Println(line1)
	fmt.Println(len(line2), line2)
//	log.Println(line2)
	t := time.Now()
	resp, err := http.PostForm("http://explain.cc:8888/app/", url.Values{"l1": {line1}, "l2": {line2}, "t": {t.Format("20060102150405")}})
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
	var line1 = ""
	var line2 = ""
	var counter = 0
	var printed = false
	var secondOrSingleLine = false

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
				if (!printed && !secondOrSingleLine) {
					httpPost(strings.TrimSpace(line1), strings.TrimSpace(line2))
					printed = true
					line1 = ""
					line2 = ""
				}
				text = false
			}

			if text {
				if secondOrSingleLine {
					line1 += translate(b)
				} else {
					line2 += translate(b)
				}


			}

			if b == 0x9b || b == 0x8c {
				secondOrSingleLine = (b == 0x8c)
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
