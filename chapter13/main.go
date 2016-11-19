package main

import (
	"bytes"
	"container/list"
	"crypto/sha1"
	"encoding/gob"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}

func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// send the message
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type", "text/html",
	)
	io.WriteString(res, `<DOCTYPE html>
<html>
  <head>
      <title>Hello World</title>
  </head>
  <body>
      Hello World!
  </body>
</html>`)
}

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server_again() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

func client_again() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}
}

func main() {
	fmt.Println(strings.Contains("test", "test"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("test", "e"),
		strings.Join([]string{"a", "b"}, "-"),
		strings.Repeat("a", 5),
		strings.Replace("aaaa", "a", "b", 2),
		strings.Split("a-b-c-d-e", "-"),
		strings.ToLower("TEST"),
		strings.ToUpper("test"),
	)
	//converting strings to bytes and vice-versa
	arr := []byte("test")
	fmt.Println(arr)
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)

	// to read or write to a []byte or string you can use the Buffer struct

	var buf bytes.Buffer
	buf.Write([]byte("test"))
	//open the file
	file, err := os.Open("./main.go")
	if err != nil {
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// create a byte buffer to read the contents in
	byte_string := make([]byte, stat.Size())
	_, err = file.Read(byte_string)
	if err != nil {
		return
	}

	str = string(byte_string)
	fmt.Println(str)
	fmt.Println("./n /n")
	fmt.Println("File size:", stat.Size()/1024, "KB.")

	// simpler way of doing it

	byte_string_again, err := ioutil.ReadFile("./main.go")
	if err != nil {
		return
	}

	str = string(byte_string_again)
	fmt.Println(str)

	// how to create a file
	new_file, err := os.Create("./test.txt")
	if err != nil {
		return
	}
	defer new_file.Close()

	new_file.WriteString("test")

	// read a directory

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	fileInfos, err := dir.Readdir(-1)
	fmt.Println(reflect.TypeOf(dir).Kind())
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		fmt.Println(info.Size())
		return nil
	})

	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	var y = list.New()
	y.PushBack(4)
	y.PushBack(5)
	y.PushBack(6)

	for e := y.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids)
	sort.Sort(ByAge(kids))
	fmt.Println(kids)

	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)

	h1, err := getHash("./main.go")
	if err != nil {
		return
	}

	h2, err := getHash("./test.txt")
	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)
	h3 := sha1.New()
	h3.Write([]byte("test"))
	bs := h3.Sum([]byte{})
	fmt.Println(bs)
	// go server()
	// go client()

	// http.HandleFunc("/hello", hello)
	// http.ListenAndServe(":9000", nil)
	// go server_again()
	// go client_again()

	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second * 1)
			fmt.Println(i, "stop")
			m.Unlock()
		}(i)
	}
	var input string
	fmt.Scanln(&input)
}
