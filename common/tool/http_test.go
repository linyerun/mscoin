package tool

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestHttpPost(t *testing.T) {
	type User struct {
		Age int `json:"age"`
	}
	http.HandleFunc("/a", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, "请求来了")
	})
	http.HandleFunc("/b", func(writer http.ResponseWriter, request *http.Request) {
		u := new(User)
		err := json.NewDecoder(request.Body).Decode(u)
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Printf("%s user=%+v\n", request.Method, u)
	})
	http.HandleFunc("/c", func(writer http.ResponseWriter, request *http.Request) {
		u := new(User)
		u.Age = 1000
		writer.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(writer).Encode(u); err != nil {
			log.Print(err)
			return
		}
	})
	http.HandleFunc("/d", func(writer http.ResponseWriter, request *http.Request) {
		u := new(User)
		err := json.NewDecoder(request.Body).Decode(u)
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Printf("%s user=%+v\n", request.Method, u)

		writer.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(writer).Encode(u); err != nil {
			log.Print(err)
			return
		}
	})
	go http.ListenAndServe(":10101", nil)
	time.Sleep(1 * time.Second)

	err := HttpPost("http://127.0.0.1:10101/a", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = HttpPost("http://127.0.0.1:10101/b", &User{Age: 200}, nil)
	if err != nil {
		log.Fatal(err)
	}

	user := new(User)
	err = HttpPost("http://127.0.0.1:10101/c", nil, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s user=%+v\n", "/c请求", user)

	err = HttpPost("http://127.0.0.1:10101/d", &User{Age: 20000}, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s user=%+v\n", "/d请求", user)
}

func TestReflectValueIsNil(t *testing.T) {
	var req any = ([]int)(nil)
	fmt.Println(reflect.ValueOf(req).IsNil())
}
