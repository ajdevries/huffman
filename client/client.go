package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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

type NodeResult struct {
	node *tree.Node
	err  error
}

func (n *Node) String() string {
	b, _ := json.Marshal(n)
	return string(b)
}

func main() {
	start := time.Now()
	flag.Parse()
	ch := make(chan *NodeResult)

	go download(*serverAddr+Huffman, "", ch)

	r := <-ch

	if r.err != nil {
		log.Fatalf("Couldn't download tree :: %q", r.err)
	}
	log.Printf("Huffman tree received :: %s\n", r.node)
	c, err := downloadCode(*serverAddr + Code)
	if err != nil {
		log.Fatalf("Couldn't download code :: %q", err)
	}
	log.Printf("Code received :: '%s'\n", c)
	d, err := r.node.Decode(c)
	if err != nil {
		log.Fatalf("Couldn't decode code :: %q", err)
	}
	log.Printf("Code decoded :: '%s'\n", d)
	log.Printf("Done :: %s\n", time.Now().Sub(start))
}

// Recursively download all nodes, and build the tree
func download(url, id string, ch chan *NodeResult) {
	u := url
	if id != "" {
		u += "/" + id
	}
	log.Printf("Downloading node :: %s\n", u)
	res, err := http.Get(u)
	if err != nil {
		ch <- &NodeResult{node: nil, err: err}
	}
	defer res.Body.Close()
	n, err := parse(res.Body)
	if err != nil {
		ch <- &NodeResult{node: nil, err: err}
	}
	log.Printf("Got node :: %s\n", n)
	node := toTreeNode(n)

	ch1 := make(chan *NodeResult)
	ch2 := make(chan *NodeResult)

	if n.Left != "" {
		go func() {
			download(url, n.Left, ch1)
		}()
	} else {
		close(ch1)
	}

	if n.Right != "" {
		go func() {
			download(url, n.Right, ch2)
		}()
	} else {
		close(ch2)
	}

	l := <-ch1
	if l != nil {
		node.Left = l.node
	}
	r := <-ch2
	if r != nil {
		node.Right = r.node
	}
	ch <- &NodeResult{node: node, err: nil}
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
