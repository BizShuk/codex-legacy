package main

import (
	"fmt"
	"sort"
)

func main() {
	input := "fdjsfkldjsf;asdjfdasopfjaoiewf"
	fmt.Printf("%s\n", input)
	hfc := huffmancode{input: input}

	hfc.char_count()
	hfc.create_tree()
	hfc.getmapcode(hfc.root, "")
	hfc.encrypt()
	hfc.diff_size()
}

type huffmancode struct {
	input      string
	ccount     map[string]int
	root       *node
	mapcode    map[string]string
	encryption string
}

func (this *huffmancode) char_count() map[string]int {
	this.ccount = make(map[string]int, len(this.input))
	for i := 0; i < len(this.input); i++ {
		this.ccount[string(this.input[i])] += 1
	}
	return this.ccount
}

func create_node(ccount map[string]int) nodelist {
	node_list := nodelist{}
	for k, v := range ccount {
		node_list = append(node_list, node{key: k, weight: v})
	}
	sort.Sort(node_list)
	return node_list
}

/*
	a1 : pick two lowest build a root node , and insert the node to sorted array
	a2 : insert merged node to another array with sort and compare two first elements of each one and pick two lowest
	other : merge with different tree side

*/
func (this *huffmancode) create_tree() {
	nodelist := create_node(this.ccount)

	for {
		if len(nodelist) == 1 {
			break
		}
		n1, n2 := nodelist[0], nodelist[1]
		nodelist = nodelist[2:]
		new := node{weight: n1.weight + n2.weight}
		new.left = &n1
		new.right = &n2
		nodelist = append(nodelist, new)
	}

	this.root = &nodelist[0]
}

func (this *huffmancode) getmapcode(root *node, code string) {
	if root.is_leaf() {
		if this.mapcode == nil {
			this.mapcode = map[string]string{}
		}
		this.mapcode[root.key] = code
		fmt.Printf("%s %s\n", root.key, code)
	}

	if root.left != nil {
		this.getmapcode(root.left, code+"0")
	}
	if root.right != nil {
		this.getmapcode(root.right, code+"1")
	}

}

func (this *huffmancode) encrypt() string {
	for i := 0; i < len(this.input); i++ {
		this.encryption += this.mapcode[string(this.input[i])]
	}
	fmt.Printf("%s\n", this.encryption)
	return this.encryption
}

func (this *huffmancode) diff_size() {
	cpercent := float32(len(this.encryption) / len(this.input))
	fmt.Printf("Comprass percentage: %.2f%", cpercent)
}

type nodelist []node

func (this nodelist) insert_node(n node) {
	this = append(this, n)
	// sort
	for i := len(this) - 1; i > 0; i-- {
		if this[i].weight < this[i-1].weight {
			this[i], this[i-1] = this[i-1], this[i]
		}
	}

}

func (this nodelist) Less(i, j int) bool {
	return this[i].weight < this[j].weight
}
func (this nodelist) Len() int {
	return len(this)
}
func (this nodelist) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type node struct {
	key         string
	weight      int
	left, right *node
}

func (this *node) is_leaf() bool {
	if this.left == nil && this.right == nil {
		return true
	}
	return false
}
