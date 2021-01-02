// demonstration of embedding of anonymous fields
package main
import "fmt"
type Point struct{X, Y int}

// anonymous fields
type Circle struct {
    Point  // anonymous field
    Radius int
}

type Wheel struct {
    Circle // anonymous field
    Spokes int
}

func main(){
	// declaration examples for anonymous fields
	// two anonymous fields of same type is not allowed
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
	Circle: Circle{
	Point: Point{X: 8, Y: 8},
	Radius: 5,
	},
	Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}
	// the adverb # prints the field names

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}
	w.X = 42
	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:42, Y:8}, Radius:5}, Spokes:20}	
}