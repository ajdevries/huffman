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

// Request struct for downloading a tree.Node based on id
type Request struct {
	id string
}

// Response struct when a Node is retrieved from the server.
type Response struct {
	id   string
	node *Node
	err  error
}

func (n *Node) String() string {
	b, _ := json.Marshal(n)
	return string(b)
}

func main() {
	start := time.Now()
	flag.Parse()
	response := make(chan *Response)
	request := make(chan *Request)

	defer func() {
		close(request)
		close(response)
	}()

	// start a go routing for downloading nodes
	go download(*serverAddr+Huffman, request, response)

	t, err := buildTree(request, response)
	if err != nil {
		log.Fatalf("Couldn't download tree :: %q", err)
	}
	log.Printf("Huffman tree received :: %s\n", t)
	c, err := downloadCode(*serverAddr + Code)
	if err != nil {
		log.Fatalf("Couldn't download code :: %q", err)
	}
	log.Printf("Code received :: '%s'\n", c)
	d, err := t.Decode(c)
	if err != nil {
		log.Fatalf("Couldn't decode code :: %q", err)
	}
	log.Printf("Code decoded :: '%s'\n", d)
	log.Printf("Done :: %s\n", time.Now().Sub(start))
}

// Start a range on the request channel for downloading nodes, send results into
// response channel.
func download(url string, request chan *Request, response chan *Response) {
	for r := range request {
		u := url
		if r.id != "" {
			u += "/" + r.id
		}
		log.Printf("Downloading node :: %s\n", u)
		res, err := http.Get(u)
		if err != nil {
			response <- &Response{err: err}
			return
		}
		defer res.Body.Close()
		n, err := parse(res.Body)
		if err != nil {
			response <- &Response{err: err}
			return
		}
		log.Printf("Got node :: %s\n", n)
		response <- &Response{id: r.id, node: n}
	}
}

// buildTree sends id's to the request channel and waits for responses on the
// response channel.
func buildTree(request chan *Request, response chan *Response) (*tree.Node, error) {
	done := make(chan bool)
	m := make(map[string]*tree.Node)
	root := &tree.Node{}
	m[""] = root
	// start with an empty id, the root node
	request <- &Request{""}
	for len(m) > 0 {

		res := <-response
		if res.err != nil {
			close(done)
			return nil, res.err
		}
		node := m[res.id]
		delete(m, res.id)
		node.Value, node.ID = res.node.Value, res.node.ID

		if res.node.Left != "" {
			m[res.node.Left] = &tree.Node{}
			node.Left = m[res.node.Left]
			go func() {
				select {
				case <-done:
				case request <- &Request{res.node.Left}:
				}
			}()
		}
		if res.node.Right != "" {
			m[res.node.Right] = &tree.Node{}
			node.Right = m[res.node.Right]
			go func() {
				select {
				case <-done:
				case request <- &Request{res.node.Right}:
				}
			}()
		}
	}
	close(done)
	return root, nil
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
