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
	// Подключение к сетевой службе.
	conn, err := net.Dial(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	// Не забываем закрыть ресурс.
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
			fmt.Printf("Найдена поговорка № %d: %s\n", id, str)
		}
	}()

	fmt.Println("Для выхода программы введите: Выход")
	s := ""
	for {
		_, err := fmt.Scanln(&s)
		if err != nil {
			return
		}
		switch s {
		case "Выход":
			log.Println("Выход из программы.")
			return
		}
	}

}
