// package main

// func main() {
// 	li, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	defer li.Close()

// 	for {
// 		con, err := li.Accept()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		go handle(conn)
// 	}
// }

// func handle(conn net.Conn) {
// 	scanner := bufio.NewScanner(conn)
// 	for scanner.Scan() {
// 		ln := scanner.Text()
// 		fmt.Println(ln)
// 	}
// 	defer conn.Close()
// }