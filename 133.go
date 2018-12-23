package sort

import (
	"sort"
	"fmt"
)

type reverse struct{ Interface }

func (r reverse) Less(i, j int) bool   { return r.Interface.Less(j, i) }
func Reverse(data Interface) Interface { return reverse{data} }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[i].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x cusromSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

sort.Sort(customSort{track,func(x,y *Track)bool{
	if x.Title!=y.Title{
		return x.Title<y.Title
	}
	if x.Year!=y.Year{
		return x.Year<y.Year
	}
	if x.Length<y.length
	}
	return false
valus:=[]int{3,1,4,1}
fmt.Println(sort.IntsAreSortted(values))
sort.Ints(values)
fmt.Println(values)
fmt.Println(sort.Reverse(sort.IntSlice(values)))
fmt.Println(values)
fmt.Println(sort.IntsAreSorted(values))
}})