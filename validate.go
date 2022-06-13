package btree

import "log"

func (bt *Btree) Validate() bool {
	var arr = []*Btree{bt}
	dupl := make(map[int]struct{})
	for {
		if len(arr) == 0 {
			return true
		}
		cur := arr[0]
		if _, exists := dupl[cur.Value]; exists {
			return false
		}
		dupl[cur.Value] = struct{}{}
		exist := bt.Contains(cur.Value)
		if exist == nil {
			return false
		}
		if exist != cur {
			return false
		}
		log.Printf("found val: %+v, cur: %+v, exist: %+v, pointers: %p %p\n", cur.Value, cur, exist, cur, exist)
		exist.Print()
		//log.Printf("exist: %+v\n", exist)
		arr = arr[1:]
		if cur.Left != nil {
			arr = append(arr, cur.Left)
		}
		if cur.Right != nil {
			arr = append(arr, cur.Right)
		}
	}
}


func (bt *Btree) NoDuplicates() bool {
	var arr = []*Btree{bt}
	found := make(map[int]struct{})
	for {
		if len(arr) == 0 {
			return true
		}
		cur := arr[0]
		if _, exists := found[cur.Value]; exists {
			return false
		}
		found[cur.Value] = struct{}{}
		arr = arr[1:]
		if cur.Left != nil {
			arr = append(arr, cur.Left)
		}
		if cur.Right != nil {
			arr = append(arr, cur.Right)
		}
	}
}

