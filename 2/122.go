var x, y IntSet
x.Add(1)
x.Add(144)
x.Add(9)
fmt.println(x.String())
y.Add(9)

y.Add(42)

fmt.println(y.String())
x.UnioWith(&y)
fmt.println(x.String())
fmt.println(x.Has(9), x.Has(123))

fmt.println(&x)
fmt.println(x.String())
fmt.println(x)


func (*IntSet)Len()int
func (*IntSet)Remove(x int)
func (*IntSet)Clear()
func (*IntSet)Copy() *IntSet