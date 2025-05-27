package main

import (
	"fmt"
	"gee_rpc"
	"log"
	"net"
	"sync"
	"time"
)

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network err : ", err.Error())
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	gee_rpc.Accept(l)
}

// v1
// func main() {
// 	addr := make(chan string)
// 	go startServer(addr)

// 	conn, _ := net.Dial("tcp", <-addr)
// 	defer func() {
// 		_ = conn.Close()
// 	}()

// 	time.Sleep(time.Second)

// 	_ = json.NewEncoder(conn).Encode(gee_rpc.DefaultOption)
// 	cc := codec.NewGobCodec(conn)

// 	for i := 0; i < 5; i++ {
// 		h := &codec.Header{
// 			ServiceMethod: "Foo.Sum",
// 			Seq:           uint64(i),
// 		}

// 		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))

// 		_ = cc.ReadHeader(h)

// 		var reply string

// 		_ = cc.ReadBody(&reply)
// 		log.Println("reply:", reply)
// 	}
// }

// v2
func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)
	client, _ := gee_rpc.Dial("tcp", <-addr)
	defer func() {
		_ = client.Close()
	}()

	time.Sleep(time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum err:", err.Error())
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
