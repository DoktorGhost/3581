package main

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	"myapp/pcg/parse"
)

// Сетевой адрес.

const addr = "0.0.0.0:12345"

// Протокол сетевой службы.
const proto = "tcp4"

func main() {
	// Запуск сетевой службы по протоколу TCP
	// на порту 12345.
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	var arr []string
	//Спарсенный срез с пословицами
	arr = parse.Parse()

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		//канал для завершения горутины
		done := make(chan bool)
		go handleConn(conn, arr, done)
	}
}

// Обработчик. Вызывается для каждого соединения.
func handleConn(conn net.Conn, arr []string, done chan bool) {
	defer conn.Close()

	go func() {
		reader := bufio.NewReader(conn)
		for {
			b, err := reader.ReadBytes('\n')
			if err != nil {
				log.Println(err)
				return
			}
			msg := strings.TrimSpace(string(b))
			// если считали не пустое сообщение от клиента - в канал отправляется true
			if msg != "" {
				done <- true
				return
			}
		}
	}()

	for {
		select {
		//если канал не пустой - завершаем цикл
		case <-done:
			return
		default:
			randomIndex := rand.Intn(len(arr)) //возможно рандомайзер нужно вынести в отдельную горутину
			message := arr[randomIndex] + "\n" // в telnet сообщения идут лесенкой, не разобрался, как это исправить, можно подсказку?
			conn.Write([]byte(message))
			time.Sleep(3 * time.Second)
		}
	}
}
