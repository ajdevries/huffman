package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ajdevries/huffman/tree"
)

func TestDownloadCode(t *testing.T) {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	mux.HandleFunc(Code, codeHandler)
	res, err := http.Get(ts.URL + Code)
	if err != nil {
		t.Fatal(err)
	}
	buf, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	expected := string(buf)
	if expected != "1000011001011001111000101000111101101001000100000000001010110111001010001111100001011" {
		t.Fatalf("Expecting a different result, got '%s'", expected)
	}
}

func TestNodeHandler(t *testing.T) {
	huffman = tree.Huffman("A", (&tree.ID{}).Next)
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	mux.HandleFunc(Huffman, nodeHandler)
	res, err := http.Get(ts.URL + Huffman)
	if err != nil {
		t.Fatal(err)
	}
	buf, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if string(buf) != "{\"ID\":\"1\",\"Value\":\"A\"}" {
		t.Fatal("Expecting a different result")
	}

	res, err = http.Get(ts.URL + Huffman + "10")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 404 {
		t.Fatal("Expecting a NotFound")
	}
}
