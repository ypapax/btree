package btree

import (
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
	t.Log(tree.String())
	as.Equal(expTree, tree)
}
