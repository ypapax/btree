package btree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {
	tree := Random(10, 100)
	tree.Print()
}

func TestPrint(t *testing.T) {
	tree := New(7).Add(77).Add(5).Add(8)//.Add(3).Add(1).Add(10).Add(10)
	tree.Print()
	//log.Println(tree.String())
}

func TestCreate(t *testing.T) {
	tree := Create(19, 7, 28, 93, 31, 80, 73, 90, 38)
	tree.Print()
	//log.Println(tree.String())
}

func TestAdd(t *testing.T) {
	as := assert.New(t)
	tree := New(7).Add(5).Add(8).Add(3).Add(1).Add(10).Add(10)
	expStr := `{
           "right/": {
              "right/": {
                 "[": 10
              },
              "[": 8
           },
           "[": 7,
           "left-_": {
              "[": 5,
              "left-_": {
                 "[": 3,
                 "left-_": {
                    "[": 1
                 }
              }
           }
        }`
	expTree, err := Parse(expStr)
	if !as.NoError(err) {
		return
	}
	tree.Print()
	//t.Log(tree.String())
	as.Equal(expTree, tree)
}


func TestValidate(t *testing.T) {
	type testCase struct{
		expectedValid bool
		tree Btree
	}
	as0 := assert.New(t)
	notValid := *Create(1,3,89,9,3989,78,909, 22323)
	node909 := notValid.Contains(909)
	node9 := notValid.Contains(9)
	//node909.Left = &Btree{Value: 908}
	if !as0.NotNil(node909) {
		return
	}
	if !as0.NotNil(node9) {
		return
	}
	node9.Left = node909

	cases := []testCase{
		//{false, Btree{Value: 10, Right: &Btree{Value: 6}}},
		//{true, Btree{Value: 10, Left: &Btree{Value: 6}}},
		//{false, Btree{Value: 6, Left: &Btree{Value: 10}}},
		//{true, *Create(1,3,89,9,3989,9)},
		//{false, Btree{Value: 10, Left: &Btree{Value: 6, Left: &Btree{Value: 10}}}},
		{false, notValid},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			as := assert.New(t)
			if !as.Equal(c.expectedValid, c.tree.Validate()) {
				c.tree.Print()
				return
			}
		})

	}

}