package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ajdevries/huffman/tree"
)

const (
	Code    = "/code/"
	Huffman = "/tree/"
)

var (
	code    = "Thanks for all the fish"
	huffman = tree.Huffman(code, ((&tree.ID{}).Next))
)

func main() {
	http.HandleFunc(Code, codeHandler)
	http.HandleFunc(Huffman, nodeHandler)
	log.Print("Listening on http://localhost:9000")
	log.Print(http.ListenAndServe(":9000", nil))
}

func codeHandler(w http.ResponseWriter, r *http.Request) {
	c, _ := huffman.Encode(code)
	fmt.Fprint(w, c)
}

func nodeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(Huffman):]
	n := huffman
	if len(id) > 0 {
		f := func(n *tree.Node) bool { return n.ID == id }
		n = huffman.FindNode(f)
		if n == nil {
			http.NotFound(w, r)
			return
		}
	}
	w.Write(n.ToJSON())
}
