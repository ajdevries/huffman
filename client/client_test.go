package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const JSON = "{\"ID\":\"%s\",\"Value\":\"A\",\"Left\":\"%s\",\"Right\":\"%s\"}"

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
	n, _ := download(ts.URL+"/tree", "")
	if n == nil {
		t.Fatal("Expecting a Node")
	}
	if n.ID != "1" {
		t.Fatal("Expecting ID == '1'")
	}
}

func TestDownloadRecursive(t *testing.T) {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	mux.HandleFunc("/tree", node("1", "2", "3"))
	mux.HandleFunc("/tree/2", node("2", "4", ""))
	mux.HandleFunc("/tree/3", node("3", "", ""))
	mux.HandleFunc("/tree/4", node("4", "", ""))
	n, _ := download(ts.URL+"/tree", "")
	if n == nil {
		t.Fatal("Expecting a Node")
	}
	if n.Left == nil || n.Left.ID != "2" {
		t.Fatal("Expecting a Left node")
	}
	if n.Right == nil || n.Right.ID != "3" {
		t.Fatal("Expecting a Right node")
	}
	if n.Left.Left == nil || n.Left.Left.ID != "4" {
		t.Fatal("Expecting a Left.Left node")
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
