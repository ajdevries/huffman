package tree

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

const (
	Left  = "1"
	Right = "0"
)

type Next func() string

type Node struct {
	ID     string
	Left   *Node
	Right  *Node
	Value  string
	Weight int
}

type Nodes []*Node

type ID struct {
	id int
}

func (id *ID) Next() string {
	id.id += 1
	return strconv.Itoa(id.id)
}

// New node, with optional to subnodes (left and right)
func New(id string, l, r *Node) *Node {
	w := 0
	if l != nil {
		w += l.Weight
	}
	if r != nil {
		w += r.Weight
	}
	return &Node{ID: id, Left: l, Right: r, Weight: w}
}

// NewLeaf build a new leaf with the given value and weight.
func NewLeaf(id, v string, w int) *Node {
	return &Node{ID: id, Value: v, Weight: w}
}

// String string representation of the tree (recursive)
func (n *Node) String() string {
	s := ""
	if n.Left != nil {
		s += fmt.Sprintf("L:%s", n.Left)
	}
	if n.Right != nil {
		s += fmt.Sprintf("R:%s", n.Right)
	}
	if n.Value != "" {
		s += fmt.Sprintf("%s", n.Value)
	}
	return s
}

// Find value in tree based on huffman encoding, returning a binary value
// as a string (for readability). When going left '1' is used, right '0'
func (n *Node) Find(v string) (string, error) {
	return n.find(v, "")
}

// Decode a binary code (represented as a string) to a value based on
// the given nodes in the three.
func (n *Node) Decode(b string) (r string, err error) {
	var v string
	for len(b) > 0 {
		v, b, err = n.decode(b)
		r += v
	}
	return
}

// Encode a value based on the given tree, returning a binary value
// represented as a string
func (n *Node) Encode(v string) (b string, err error) {
	var r string
	for _, val := range []byte(v) {
		r, err = n.Find(string(val))
		if err != nil {
			return
		}
		b += r
	}
	return
}

func (n *Node) decode(b string) (string, string, error) {
	if n.IsLeaf() {
		return n.Value, b, nil
	}
	if len(b) > 0 {
		if string(b[0]) == Left {
			return n.Left.decode(string(b[1:]))
		}
		if string(b[0]) == Right {
			return n.Right.decode(string(b[1:]))
		}
	}
	return "", "", errors.New("Broken encoded value")
}

// IsLeaf is this node a leaf (no left or right nodes attached)
func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) find(v, path string) (r string, err error) {
	if n.Value == v {
		return path, nil
	}
	if n.IsLeaf() {
		return "", fmt.Errorf("Value %s not found in tree", v)
	}
	if n.Left != nil {
		r, err = n.Left.find(v, path+Left)
	}
	if err != nil && n.Right != nil {
		r, err = n.Right.find(v, path+Right)
	}
	return r, err
}

func (s Nodes) Len() int {
	return len(s)
}

func (s Nodes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Nodes) Less(i, j int) bool {
	if s[i].Weight == s[j].Weight {
		return s[i].Value < s[j].Value
	}
	return s[i].Weight > s[j].Weight
}

// Huffman building an huffman tree based on chars from a string
func Huffman(value string, n Next) *Node {
	l := weightedNodes(value, n)
	var left, right *Node
	for len(l) > 1 {
		sort.Sort(l)
		left, l = l[len(l)-1], l[:len(l)-1]
		right, l = l[len(l)-1], l[:len(l)-1]
		l = append(l, New(n(), left, right))
	}
	return l[0]
}

func weightedNodes(value string, n Next) Nodes {
	l := Nodes{}
	w := map[string]int{}
	for _, v := range []byte(value) {
		val := string(v)
		if f, ok := w[val]; ok {
			w[val] = f + 1
		} else {
			w[val] = 1
		}
	}
	for k, v := range w {
		l = append(l, NewLeaf(n(), k, v))
	}
	return l
}
