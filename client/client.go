package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ajdevries/huffman/tree"
)

const (
	// Code endpoint
	Code = "/code"
	// Huffman endpoint
	Huffman = "/tree"
)

var (
	serverAddr = flag.String("server", "http://localhost:9000", "Address and port from the server")
)

// Node representation in the client, used to deserialize the tree.Node from
// the server.
type Node struct {
	ID    string
	Left  string
	Right string
	Value string
}

func (n *Node) String() string {
	b, _ := json.Marshal(n)
	return string(b)
}

func main() {
	flag.Parse()
	huffman, err := download(*serverAddr+Huffman, "")
	if err != nil {
		log.Fatalf("Couldn't download tree :: %q", err)
	}
	log.Printf("Huffman tree received :: %s\n", huffman)
	c, err := downloadCode(*serverAddr + Code)
	if err != nil {
		log.Fatalf("Couldn't download code :: %q", err)
	}
	log.Printf("Code received :: '%s'\n", c)
	d, err := huffman.Decode(c)
	if err != nil {
		log.Fatalf("Couldn't decode code :: %q", err)
	}
	log.Printf("Code decoded :: '%s'\n", d)
}

// Recursively download all nodes, and build the tree
func download(url, id string) (*tree.Node, error) {
	u := url
	if id != "" {
		u += "/" + id
	}
	log.Printf("Downloading node :: %s\n", u)
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	n, err := parse(res.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("Got node :: %s\n", n)
	node := toTreeNode(n)
	if n.Left != "" {
		left, err := download(url, n.Left)
		if err != nil {
			return nil, err
		}
		node.Left = left
	}
	if n.Right != "" {
		right, err := download(url, n.Right)
		if err != nil {
			return nil, err
		}
		node.Right = right
	}
	return node, nil
}

// Parse the node info (as JSON) from the server
func parse(r io.Reader) (*Node, error) {
	n := &Node{}
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func downloadCode(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func toTreeNode(n *Node) *tree.Node {
	return &tree.Node{ID: n.ID, Value: n.Value}
}
