package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

const addr = "localhost:12345"
const network = "tcp4"
const link = "https://go-proverbs.github.io/"

func main() {
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			return
		}
	}(listener)

	response, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	tkn := html.NewTokenizer(strings.NewReader(string(body)))
	var phrases []string
	var h3, a bool
	for {
		nextTkn := tkn.Next()
		if nextTkn == html.ErrorToken {
			break
		} else if nextTkn == html.StartTagToken {
			tknData := tkn.Token()
			if tknData.Data == "h3" {
				h3 = true
			} else if tknData.Data == "a" {
				a = true
			}
		} else if nextTkn == html.TextToken {
			if h3 && a {
				tknData := tkn.Token()
				phrases = append(phrases, tknData.Data)
				h3 = false
				a = false
			}
		}
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go handleConn(conn, phrases)
		}
	}()

	fmt.Println("Введите exit для выхода из программы")
	s := ""
	for {
		_, err := fmt.Scanln(&s)
		if err != nil {
			return
		}
		switch s {
		case "exit":
			log.Println("Завершение работы программы")
			return
		}
	}

}

func handleConn(conn net.Conn, phrases []string) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	for {
		_, err := conn.Write([]byte(phrases[rand.Intn(len(phrases))] + "\n\r"))
		if err != nil {
			return
		}

		time.Sleep(3 * time.Second)
	}
}
