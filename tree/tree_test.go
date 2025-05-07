package tree

import (
	"encoding/json"
	"testing"
)

type MyNode struct {
	ID  int64 `json:"id"`
	PID int64 `json:"pid"`
}

func (n *MyNode) GetID() int64 {
	return n.ID
}

func (n *MyNode) GetParentID() int64 {
	return n.PID
}

func (n *MyNode) GetOrder() int {
	return -1
}

// 随机生成树形结构, id大小随机
// pid随机
var data = `
[
	{"id":1,"pid":0},
	{"id":2,"pid":3},
	{"id":3,"pid":1},
	{"id":4,"pid":1},
	{"id":5,"pid":2},
	{"id":6,"pid":2},
	{"id":7,"pid":4},
	{"id":8,"pid":4},
	{"id":9,"pid":5},
	{"id":10,"pid":5},
	{"id":11,"pid":6},
	{"id":12,"pid":6},
]
`

func TestCase1(t *testing.T) {
	var nodes []MyNode
	if err := json.Unmarshal([]byte(data), &nodes); err != nil {
		t.Fatal(err)
	}

	tree := BuildTree(nodes)
	if tree == nil {
	}
}
