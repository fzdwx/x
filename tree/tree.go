package tree

import (
	"sort"
)

type Tree struct {
	Root  Node
	Nodes []Node
}

type Node interface {
	GetID() int64
	GetParentID() int64
	GetOrder() int
}

// BuildTree 根据节点列表构建树形结构
// nodes: 节点列表
// 如果节点列表为空，则返回 nil
// 如果pid为0，则表示根节点, 最后如果没有找到根节点，则第一个节点为根节点
func BuildTree[T Node](nodes []T) *Tree {
	if len(nodes) == 0 {
		return nil
	}

	// 转换为Node接口切片，方便在返回的Tree中统一处理
	nodeInterfaces := make([]Node, len(nodes))
	for i := range nodes {
		nodeInterfaces[i] = nodes[i]
	}

	nodesMap := make(map[int64]Node)
	for _, node := range nodeInterfaces {
		nodesMap[node.GetID()] = node
	}

	var root Node
	childrenMap := make(map[int64][]Node)
	for _, node := range nodeInterfaces {
		if node.GetParentID() == 0 {
			root = node
		} else {
			childrenMap[node.GetParentID()] = append(childrenMap[node.GetParentID()], node)
		}
	}

	if root == nil && len(nodeInterfaces) > 0 {
		root = nodeInterfaces[0]
	}

	return buildTreeRecursion(root, childrenMap)
}

// Traversal 遍历树形结构
func Traversal(root *Tree, fn func(node Node)) {
	_ = TraversalE(root, func(node Node) error {
		fn(node)
		return nil
	})
}

// TraversalE 遍历树形结构，支持错误处理
func TraversalE(root *Tree, fn func(node Node) error) error {
	if root == nil {
		return nil
	}

	if err := fn(root.Root); err != nil {
		return err
	}
	for _, node := range root.Nodes {
		if node != root.Root {
			if err := fn(node); err != nil {
				return err
			}
		}
	}
	return nil
}

func buildTreeRecursion(root Node, childrenMap map[int64][]Node) *Tree {
	if root == nil {
		return nil
	}

	node := &Tree{Root: root, Nodes: []Node{root}}

	if children, ok := childrenMap[root.GetID()]; ok {
		sort.Slice(children, func(i, j int) bool {
			return children[i].GetOrder() < children[j].GetOrder()
		})
		for _, child := range children {
			childTree := buildTreeRecursion(child, childrenMap)
			if childTree != nil {
				node.Nodes = append(node.Nodes, childTree.Nodes...)
			}
		}
	}

	return node
}
