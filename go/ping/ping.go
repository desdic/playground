package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type IOPingResult struct {
	MS  float64
	Err error
}

func IOPing(filename string, filesize int, waittime int) <-chan IOPingResult {
	r := make(chan IOPingResult)

	generateData := func(n int) string {
		letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b := make([]rune, n)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return string(b)
	}

	go func() {
		data := generateData(filesize)
		f, err := os.Create(filename)
		if err != nil {
			r <- IOPingResult{0.0, err}
			return
		}
		w := bufio.NewWriter(f)
		_, err = w.WriteString(data)
		if err != nil {
			r <- IOPingResult{0.0, err}
			return
		}
		w.Flush()
		f.Close()

		buf := make([]byte, filesize)

		for {
			start := time.Now()

			f, err = os.Open(filename)
			if err != nil {
				r <- IOPingResult{0.0, err}
				return
			}
			numb, err := f.Read(buf)
			if err != nil {
				r <- IOPingResult{0.0, err}
				return
			}
			if numb != filesize {
				r <- IOPingResult{0.0, fmt.Errorf("bytes read differs from bytes written")}
				return
			}
			f.Close()
			elapsed := time.Since(start)
			// in ms
			r <- IOPingResult{float64(elapsed.Microseconds()) * 0.001, nil}
			time.Sleep(time.Duration(waittime) * time.Second)
		}
	}()

	return r
}

func main() {
	pings := IOPing("./ioping", 1024*1024*10, 1)

	for res := range pings {
		if res.Err != nil {
			log.Panic(res.Err)
		}
		log.Printf("%2.2f ms", res.MS)
	}
}
