package btree

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/ypapax/jsn"
)

type Btree struct {
	Left  *Btree
	Value int
	Right *Btree
}

func init() {
	rand.Seed(time.Now().Unix())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func New(val int) *Btree {
	return &Btree{Value: val}
}

func Random(nodes, maxVal int) *Btree {
	root := &Btree{Value: int(rand.Int63n(int64(maxVal)))}
	for i := 1; i < nodes; i++ {
		val := int(rand.Int63n(int64(maxVal)))
		root.Add(val)
	}
	return root
}

func (bt *Btree) Add(val int) *Btree {
	if val <= bt.Value {
		if bt.Left == nil {
			bt.Left = &Btree{Value: val}
			return bt
		}
		bt.Left.Add(val)
		return bt
	}
	if bt.Right == nil {
		bt.Right = &Btree{Value: val}
		return bt
	}
	bt.Right.Add(val)
	return bt
}

func (bt *Btree) String() string {
	return jsn.B(bt)
}

func Parse(str string) (*Btree, error) {
	var tree Btree
	if err := json.Unmarshal([]byte(str), &tree); err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return &tree, nil
}
