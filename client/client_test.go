package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

const JSON = "{\"ID\":\"%s\",\"Value\":\"A\",\"Left\":\"%s\",\"Right\":\"%s\"}"

func TestMain(m *testing.M) {
	c := make(chan bool)
	go func() {
		os.Exit(m.Run())
		c <- true
	}()
	select {
	case <-c:
	case <-time.After(2 * time.Second):
		log.Fatal("Timeout")
	}

}

func TestParse(t *testing.T) {
	n, _ := parse(strings.NewReader(fmt.Sprintf(JSON, "1", "", "")))
	if n == nil {
		t.Fatal("Expecting a leaf node here")
	}
	if n.ID != "1" {
		t.Fatalf("Expecting ID == 1, but got %s", n.ID)
	}
	_, err := parse(strings.NewReader("No JSON"))
	if err == nil {
		t.Fatal("Expecting an error here")
	}
}

func TestToTreeNode(t *testing.T) {
	n, _ := parse(strings.NewReader(fmt.Sprintf(JSON, "1", "", "")))
	node := toTreeNode(n)
	if node == nil {
		t.Fatal("Didn't expect nil")
	}
	if node.ID != "1" {
		t.Fatal("Expecting ID 1")
	}
	if node.Value != "A" {
		t.Fatal("Expecting Value 'A'")
	}
}

func TestDownload(t *testing.T) {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	mux.HandleFunc("/tree", node("1", "", ""))
	request := make(chan *Request)
	response := make(chan *Response)

	go download(ts.URL+"/tree", request, response)
	request <- &Request{id: ""}
	res := <-response
	n := res.node
	if n == nil {
		t.Fatal("Expecting a Node")
	}
	if n.ID != "1" {
		t.Fatal("Expecting ID == '1'")
	}
}

func TestDownloadError(t *testing.T) {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	mux.HandleFunc("/tree", node("1", "", ""))
	request := make(chan *Request)
	response := make(chan *Response)

	go download(ts.URL, request, response)
	request <- &Request{id: ""}
	res := <-response

	if res.err == nil {
		t.Fatal("Expecting a error")
	}

	go download("http://localhost/tree", request, response)
	request <- &Request{id: ""}
	res = <-response

	if res.err == nil {
		t.Fatal("Expecting a error")
	}

}

func TestBuildTree(t *testing.T) {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	mux.HandleFunc("/tree", node("1", "", ""))
	request := make(chan *Request)
	response := make(chan *Response)

	go download(ts.URL+"/tree", request, response)
	tree, _ := buildTree(request, response)
	if tree == nil || tree.ID != "1" {
		t.Fatal("Expecting a tree with ID 1")
	}
}

func TestBuildTreeRecursive(t *testing.T) {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	mux.HandleFunc("/tree", node("1", "2", "3"))
	mux.HandleFunc("/tree/2", node("2", "4", ""))
	mux.HandleFunc("/tree/3", node("3", "", ""))
	mux.HandleFunc("/tree/4", node("4", "", ""))
	request := make(chan *Request)
	response := make(chan *Response)

	go download(ts.URL+"/tree", request, response)
	tree, _ := buildTree(request, response)
	if tree == nil || tree.ID != "1" {
		t.Fatal("Expecting a tree with ID 1")
	}
	if tree.Left == nil || tree.Left.ID != "2" {
		t.Fatal("Expecting a first left Node with id 2")
	}
	if tree.Left.Left == nil || tree.Left.Left.ID != "4" {
		t.Fatal("Expecting a first left Node with id 4")
	}
	if tree.String() != "L:L:AAR:AA" {
		t.Fatalf("Expecting 'L:L:AAR:AA' but got %s", tree)
	}
}

func TestBuildTreeError(t *testing.T) {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	mux.HandleFunc("/tree", node("1", "2", "3"))
	request := make(chan *Request)
	response := make(chan *Response)

	go download(ts.URL+"/tree", request, response)
	_, err := buildTree(request, response)
	if err == nil {
		t.Fatal("Expecting an error")
	}
}

func TestDownloadCode(t *testing.T) {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	mux.HandleFunc("/code", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "1234")
	})

	c, _ := downloadCode(ts.URL + Code)
	if c == "" {
		t.Fatalf("Expecting a different code, but got '%s'", c)
	}
}

func node(id, left, right string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, JSON, id, left, right)
	}
}
