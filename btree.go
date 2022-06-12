package btree

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/ypapax/jsn"
)

type Btree struct {
	Right *Btree `json:"right/,omitempty"`
	Value int `json:"["`
	Left  *Btree `json:"left-_,omitempty"`
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

func Create(vals ...int) *Btree {
	if len(vals) == 0 {
		return nil
	}
	b := New(vals[0])
	for _, v := range vals[1:] {
		b.Add(v)
	}
	return b
}

func (bt *Btree) Add(val int) *Btree {
	if bt.Value == val {
		return bt
	}
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

type printBtreeObj struct {
	Btree
	isLeft bool
	isRight bool
	Row    int
	Column int
	Empty bool
	Parent *printBtreeObj
}

func (bt *Btree) Print() {
	m := bt.GetPrintMatrix()
	PrintMatrix(m)
}

func PrintMatrix(result [][]string) {
	var maxWidth int
	firstLeftNotEmptyColumn := -1
	lastRightNotEmptyColumn := -1
	log.Println("len(result)", len(result))
	for _, row := range result {
		log.Println("len(row)", len(row))
		for j, el := range row {
			if len(strings.TrimSpace(el)) > 0 && (j < firstLeftNotEmptyColumn || firstLeftNotEmptyColumn < 0) {
				firstLeftNotEmptyColumn = j
			}
			if len(strings.TrimSpace(el)) > 0 && j > lastRightNotEmptyColumn {
				lastRightNotEmptyColumn = j
			}
			if len(el) > maxWidth {
				maxWidth = len(el)
			}
		}
	}
	log.Println("firstLeftNotEmptyColumn", firstLeftNotEmptyColumn)
	for _, row := range result {
		var line string
		for j, el := range row {
			if j < firstLeftNotEmptyColumn {
				continue
			}
			if j > lastRightNotEmptyColumn {
				continue
			}
			elSpace := addSpaces(maxWidth, el)
			line += "_"+elSpace
		}
		//log.Println("len(line): ", len(line), "len(row)", len(row), "len(result)", len(result))
		log.Println(line)
	}
}
const spaceChar = "."
func addSpaces(maxWidth int, el string) string {
	//log.Println("maxWidth", maxWidth)
	spaceLenLeft := maxWidth - len(el)
	var space string
	//for i := 0; i < int(math.Ceil(float64(spaceLenLeft)/2)); i++ {
	for i := 0; i < spaceLenLeft; i++ {
		space += spaceChar
	}
	return space + el
}

func (bt *Btree) GetPrintMatrix() (result [][]string) {
	if bt == nil {
		return nil
	}
	var byLevel = make(map[int][]printBtreeObj)
	var cur printBtreeObj
	curLevel := 0
	q := printBtreeObj{Btree: *bt}
	byLevel[curLevel] = append(byLevel[curLevel], q)
	var queue []printBtreeObj
	queue = append(queue, q)
	var maxLevel int
	var maxElementsOnRow = 1
	for {
		if len(queue) == 0 {
			break
		}
		cur = queue[0]
		queue = queue[1:]
		newLevel := cur.Row + 1
		var gotChildren bool
		if cur.Left != nil {
			ql := printBtreeObj{Btree: *cur.Left, isLeft: true, Row: newLevel, Parent: &cur, Column: cur.Column - 1}
			byLevel[newLevel] = append(byLevel[newLevel], ql)
			maxLevel = newLevel
			queue = append(queue, ql)
			gotChildren = true
		} else {
			byLevel[newLevel] = append(byLevel[newLevel], printBtreeObj{isLeft: true, Empty: true, Parent: &cur, Column: cur.Column - 1})
		}
		if cur.Right != nil {
			qr := printBtreeObj{Btree: *cur.Right, isRight: true, Row: newLevel, Parent: &cur, Column: cur.Column+1}
			byLevel[newLevel] = append(byLevel[newLevel], qr)
			maxLevel = newLevel
			queue = append(queue, qr)
			gotChildren = true
		} else {
			byLevel[newLevel] = append(byLevel[newLevel], printBtreeObj{isRight: true, Empty: true, Parent: &cur, Column: cur.Column+1})
		}
		if gotChildren {
			maxElementsOnRow *= 2
		}
	}
	maxElementsOnRowIncludingSpace := maxElementsOnRow * 2
	log.Println("maxElementsOnRow", maxElementsOnRow)
	maxElOnTheRow := 1
	log.Println("maxLevel", maxLevel)
	result = make([][]string, maxLevel+1)
	headColumn := maxElementsOnRowIncludingSpace / 2
	for level := 0; level <= maxLevel; level++ {
		result[level] = make([]string, maxElementsOnRowIncludingSpace)

		els := byLevel[level]
		spaceFloat := float64(maxElementsOnRow- maxElOnTheRow)/2
		space := /*centerColumn+*/int(math.Round(spaceFloat))
		log.Println("space", space, "els", len(els), "maxElOnTheRow", maxElOnTheRow, "spaceFloat", spaceFloat)
		var spaceLeft string
		for i := 0; i < space; i++ {
			spaceLeft += " "
		}
		var row = spaceLeft
		for j, el := range els {
			var valStr string
			/*if el.Parent != nil {
				byLevel[level][i].Column = el.Parent.Column + 1
			}*/
			if el.Btree.Value > 0 {
				valStr = fmt.Sprintf("%+v", el.Btree.Value)
			}
			column := headColumn+el.Column
			if len(result) <= level {
				panic(fmt.Sprintf("not enough lines count %+v, len(result): %+v", level, len(result)))
			}
			if len(result[level]) <= column {
				panic(fmt.Sprintf("not enough columns %+v, len(result[level]): %+v for level %+v, j: %+v, len(els): %+v, headColumn: %+v",
					column, len(result[level]), level, j, len(els), headColumn))

			}
			result[level][column] = valStr
			//row = row +  fmt.Sprintf("%+v %+v", spaceLeft, valStr)
		}
		row += spaceLeft
		log.Println(row)
		maxElOnTheRow *= 2
	}
	return result
}

