package main

import (
	"fmt"
	"strings"
	"sync"

	ocpb "github.com/openconfig/reference/rpc/openconfig"
)

var (
	root *node
)

type Root struct {
	mu sync.Mutex
	c  map[string]*node
	i  map[string]*node
}

func (r *Root) Get(p *Path) (*node, bool) {
	if p == nil {
		return nil, false
	}
	n, ok := r.i[p.Key()]

}

type node struct {
	c map[Key]*node
	p []string
	n string
	v ocpb.Value
}

func (n *node) GetChild(path []string) (*node, bool) {
	n, ok := n.c[NewPath(path)]
	return n, ok
}

type Key interface {
	Key() string
}

type Path ocpb.Path

func NewPath(elements []string) *Path {
	e := make([]string, len(elements))
	copy(elements, e)
	p := &Path{
		Element: e,
	}
	return p
}

func (p *Path) Key() string {
	if p == nil {
		return "/"
	}
	return strings.Join(p.Element, "/")
}

func main() {
	root := &node{
		n: "device",
		c: map[Key]*node{
			NewPath([]string{"protocols"}): &node{
				n: "protocols",
				c: map[Key]*node{
					NewPath([]string{"bgp"}): &node{
						n: "bgp",
					},
					NewPath([]string{"ospf"}): &node{
						n: "ospf",
					},
				},
			},
		},
	}
	fmt.Printf("%+v\n", root)
	c, _ := root.GetChild([]string{"protocols"})
	fmt.Printf("%+v\n", c)
}
