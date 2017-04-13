package tree

import (
	"fmt"
	"sort"
)

type Node struct {
	Left   *Node
	Right  *Node
	Value  string
	Weight int
}

type Nodes []*Node

func New(l, r *Node) *Node {
	w := 0
	if l != nil {
		w += l.Weight
	}
	if r != nil {
		w += r.Weight
	}
	return &Node{Left: l, Right: r, Weight: w}
}

func NewLeaf(v string, w int) *Node {
	return &Node{Value: v, Weight: w}
}

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

func (s Nodes) Len() int {
	return len(s)
}

func (s Nodes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Nodes) Less(i, j int) bool {
	return s[i].Weight > s[j].Weight
}

func Huffman(values ...string) *Node {
	l := Nodes{}
	for _, v := range values {
		l = append(l, NewLeaf(v, 1))
	}
	var left, right *Node
	for len(l) > 1 {
		left, l = l[len(l)-1], l[:len(l)-1]
		right, l = l[len(l)-1], l[:len(l)-1]
		l = append(l, New(left, right))
		sort.Sort(l)
	}
	return l[0]
}
