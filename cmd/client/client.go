package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	addr  = "localhost:12345"
	proto = "tcp4"
)

func main() {
	conn, err := net.Dial(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	reader := bufio.NewReader(conn)
	id := 0

	go func() {
		for {
			pvb, err := reader.ReadBytes('\n')
			if err != nil {
				log.Fatal(err)
			}
			id++
			str := strings.Trim(string(pvb), "\n")
			str = strings.Trim(str, "\r")
			fmt.Printf("Поговорка: %d: %s\n", id, str)
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
