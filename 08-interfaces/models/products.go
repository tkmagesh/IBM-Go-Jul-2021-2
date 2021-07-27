package models

import (
	"fmt"
	"sort"
)

type Products []Product

func (ps Products) Print() {
	for _, v := range ps {
		fmt.Println(v)
	}
}

/* implement the sort apis */

func (ps Products) Len() int           { return len(ps) }
func (ps Products) Less(i, j int) bool { return ps[i].Id < ps[j].Id }
func (ps Products) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

func (ps Products) Sort() {
	sort.Sort(ps)
}

type byName struct {
	Products
}

func (ps byName) Less(i, j int) bool {
	return ps.Products[i].Name < ps.Products[j].Name
}

func (ps Products) SortByName() {
	sort.Sort(byName{ps})
}
