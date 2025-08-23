package cache

type Btree struct {
	Root *Node
}

type Node struct {
	Key    string
	Value  any
	TTI    string
	Left   *Node
	Right  *Node
	Parent *Node
}

func (bt *Btree) Get(key string) (any, bool) {
	node := bt.Root
	for node != nil {
		if node.Key == key {
			return node.Value, true
		}
		if key < node.Key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil, false
}

func (bt *Btree) Set(key string, value any) {
	bt.Root = bt.insert(bt.Root, key, value)
}

func (bt *Btree) Delete(key string) {
	bt.Root = bt.remove(bt.Root, key)
}

func (bt *Btree) Clear() {
	bt.Root = nil
}

func (bt *Btree) ListAllNodes() []Node {
	var nodes []Node
	bt.traverse(bt.Root, func(n *Node) {
		nodes = append(nodes, *n)
	})
	return nodes
}

func (bt *Btree) RemoveNode(key string) {
	bt.Root = bt.remove(bt.Root, key)
}

func (bt *Btree) insert(node *Node, key string, value any) *Node {
	if node == nil {
		return &Node{
			Key:   key,
			Value: value,
		}
	}
	if key < node.Key {
		node.Left = bt.insert(node.Left, key, value)
	} else {
		node.Right = bt.insert(node.Right, key, value)
	}
	return node
}

func (bt *Btree) remove(node *Node, key string) *Node {
	if node == nil {
		return nil
	}
	if key < node.Key {
		node.Left = bt.remove(node.Left, key)
	} else if key > node.Key {
		node.Right = bt.remove(node.Right, key)
	} else {
		// Node to be deleted found
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		// Node has two children, find the inorder successor
		minNode := bt.findMin(node.Right)
		node.Key = minNode.Key
		node.Value = minNode.Value
		node.Right = bt.remove(node.Right, minNode.Key)
	}
	return node
}

func (bt *Btree) findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (bt *Btree) traverse(node *Node, visit func(*Node)) {
	if node == nil {
		return
	}
	bt.traverse(node.Left, visit)
	visit(node)
	bt.traverse(node.Right, visit)
}
