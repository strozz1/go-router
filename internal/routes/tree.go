package routes

import "log"

type Comparable[K any] interface {
	Equals(k K) bool
    Debug() string
}


// type def for a tree node
type Node[T Comparable[T]] struct {
	value    T
	children []*Node[T]
}

func (n *Node[T]) AddChild(child Node[T]) {
	n.children = append(n.children, &child)
}

// returns the pointer to the child if the value exists in the children
func (node *Node[T]) GetChild(finder T) *Node[T] {
	var child *Node[T] = nil
	if node == nil {
		return nil
	}
	for _, v := range node.children {
		if v.value.Equals(finder) {
			child = v
			break
		}
	}
	return child
}

// tree data structure
type Tree[T Comparable[T]] struct {
	root   *Node[T]
	height int
}

// creates an empty tree with no children
func EmptyTree[T Comparable[T]]() Tree[T] {
	return Tree[T]{
		root: &Node[T]{
			children: []*Node[T]{},
		},
		height: 1,
	}
}

// Return the node containing the value T. Nil if not found
func (t *Tree[T]) Search(value T) *Node[T] {
	return walk(t.root, value)
}

func walk[T Comparable[T]](current *Node[T], value T) *Node[T] {
	if current.value.Equals(value) {
		return current
	}
	for _, n := range current.children {
		m := walk(n, value)
		if m != nil {
			return m
		}
	}
	//if no value find
	return nil

}

func (t *Tree[T]) Print() {
	indent := ""
	printTree[T](t.root, indent)

}


func printTree[T Comparable[T]](node *Node[T],indent string ) {
	log.Printf("%s|%v", indent, node.value.Debug())
	indent = indent + "--"
	for _, n := range node.children {
		printTree[T](n, indent)
	}

}
