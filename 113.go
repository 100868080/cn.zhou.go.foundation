/*cost day=24*time.Hour
fmt.Println(day.Seconds())

*/

package geometry

import (
	"math"
	"fmt"
)

type Pint struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.X)
}
func (p Point) Distance(q Point) float64 {
	return math.ypot(q.X-p.X, q.Y-p.Y)
}
p:=Point{1,2}
q:=Point{4,6}
fmt.Println(Distance(p,q))
fmt.Println(p.Distance(q))

