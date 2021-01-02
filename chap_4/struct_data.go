// demo code for structs 
package main
import (
    "fmt"
    "time"
)

// structs is an aggregate data-type that groups together zero or mode 
// named values of arbitrary type
// Each value in a struct is called a field

// declare a struct type Employee
// same type of variables could be declared in same line
// i.e 
// Name, Postion    string

// Field order is significant to type identity
// only fields that start with caps are exported (private/public)
type Employee struct{
    ID          int
    Name        string
    DoB         time.Time
    Position    string
    Salary      int
    ManagerID   int
}
// declare variable of the type Employee
var boss Employee

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

    // fields of the instance can be accessed using dot notation 
    var dilbert Employee
    dilbert.Name = "Dilbert"
    dilbert.Salary = 12000
    dilbert.Salary -= 5000 // demoted, for writing too few lines of code

    // or, flieds can be accessed using pointers
    // assign address of dilbert.Position to a pointer
    position := &dilbert.Position
    // access the value using asterisk operator
    *position = "Senior" + *position  // promoted, for outsourcing to Elbonia
    fmt.Println(dilbert)

    // address for var dilbert can also be used directly
    // declare a pointer of the type Employee and assign it address of dilbert
    var employeeOfTheMonth *Employee = &dilbert
    // fields can be accesed using dot operator along with the pointer
    // (*employeeOfTheMonth).Position += " (proactive team player)"
    // or the statement below are equivalent
    employeeOfTheMonth.Position += " (proactive team player)"
    fmt.Println(dilbert)

    // dot notation can be directly used with struct pointers
    fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

    // the result of the func all is a pointer of type struct
    // if the function were to return a struct of type Empolyee and
    // not a pointer, this assignment would fail to. This would happen since 
    // the left hand side would not be identified as a variable
    EmployeeByID(0).Salary = 0

    // Struct Literals
    // a struct can be initialized at declaration using a struct literal
    
    // the values are assigned by the order in the list corresponding to the 
    // ones that are define in the struct
    p1 := Point{1, 2}
    // struct can also be intiialized by listing some or all of 
    // field names, along with corresponding values
    // if a value is not given, it is initliazled by default to the varaible zero value
    // in this case, the ordering of the variable do not matter
    p2 := Point{Y:1} 
    fmt.Println(p1, p2)

    fmt.Println(Scale(p1, 5))   // {5, 10}
    // however, for efficiency , larger struct types are passed and returned using pointers
    // structs can be directly dealt through pointers
    p3 := &Point{1,2}  
    // or equivalently
    p4 := new(Point)  // new assignes a pointer of the type Point
    *p4 = Point{1, 2}
    fmt.Printf("&Point : %T, new(Point) : %T\n", p3, p4)

    // If all the fields of a struct are comparable. then the struct is also comparable
    fmt.Printf("p1.X == p2.X && p1.Y == p2.Y: %v\n", p1.X == p2.X && p1.Y == p2.Y) // false
    fmt.Printf("p1 == p2 : %v\n", p1==p2) // false
    
    // comparable stucts can thus be used as key type of map
    coord := make(map[Point]int)
    coord[Point{0,0}]++
    coord[Point{0,1}]++
    fmt.Println(coord)

    // Go provides convenient chaining of feilds, such that
    // x.d.e.f can be expressed just using x.f
    // if  we had struct like

    type Circle1 struct {
        Center Point
        Radius int
    }

    type Wheel1 struct {
        Circle Circle1
        Spokes int
   }

    // the elements  would be accessed in the following fashion
    var w1 Wheel1
    _  = Wheel1{Circle1{Point{8, 8}, 5}, 20}
    w1.Circle.Center.X  = 8
    w1.Circle.Center.Y = 8
    w1.Circle.Radius = 5
    w1.Spokes = 20
    fmt.Println(w1)

    // using structs with anonymous fields, thanks to embedding, 
    // the names can be refered directly at the leaves of the implicit tree, 
    // without the intervening names
    // this only works with anonymous fields
    var w Wheel
    // however, the following will fail
    // w = Wheel {X: 8, Y: 8, Radius: 5, Spokes: 20}  // compile error : unknown fields
    // w = Wheel{8, 8, 5, 20}                         // compile error: unknown fields
    // another way of initializing
    _  = Wheel{Circle{Point{8, 8}, 5}, 20}

    w.X = 8         // equivalent to w.Circle.Point.X = 8
    w.Y = 8         // equivalent to w.Circle.Point.X = 8
    w.Radius = 5    // equivalent to w.Circle.Radius = 5
    w.Spokes = 20
    fmt.Println(w)
}


//scales a point
func Scale(p Point, factor int) Point {
    return Point{p.X * factor, p.Y *  factor}
}



// This dummy function returns an pointer to an Employee struct
func EmployeeByID(id int) *Employee {
    boss.Position = "Pointy-haired boss"
    return &boss
}
