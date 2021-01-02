
// demp codes for maps

package main
import (
	"fmt"
	"sort"
)

func main(){
    // map is a reference to a hash table
    // syntax - map[K]V where K and V are types of key and value
    // all keys are of the same type and all values are of the same type
	// type K of keys must be comparable using "=="

	
	ages := map[string]int{  // mapping from strings to ints
		"alice":   31,
		"charlie": 34,
	}
	
	// built-in make can be used to create map
	// the above code is equivalent to
	/*
	ages := make(map[string]int)
	ages["alice"]  = 31
	ages["charlie"] = 34
	*/
	// a new empty map can be also be created using
	// ages := map[string]int{}
	fmt.Println(ages)
	// map elements are accessed using the subscript notation
	ages["alice"] = 32
	fmt.Println(ages["alice"])	// 32

	// a key-value pair can be removed using built-in delete function
	delete(ages, "alice")
	fmt.Println(ages)

	// if an element is not present in the map, a lookup for that element
	// returns the zero value of the value-type
	// so, something like this is valid
	// ages["bob"] = ages["bob"] + 1 // happy birthday!
	// or, alternatively
	// ages["bob"] += 1
	// or just
	ages["bob"]++
	fmt.Println(ages)

	// map can in iterated using range-based for loop
	// successive iteration causes key, value pair to be set to the variables 
	// on the left-hand side of the range expression
	
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	// the order of map iteration is unspecified. 
	// to enumerate key/value pair in order, the keys must be sorted explicitly
	// we can declare slice of the key-type

	// var names []string
	// However, since we know the size of the map beforehand, it's more effecient to 
	// allocate array or slice of fixed size
	// initialize slice of len 0 , capacity as len(ages)
	names := make([]string, 0, len(ages))

	// if only one variable is used on left-hand side of range expression, 
	// the key (or for arrays/slice index) is updated during iteration
	for name := range ages{
		names = append(names, name)
	}
	// sort names
	sort.Strings(names)
	// since we need the elements, the first variable is left as blank
	// now we can iterate map in sorted key fashion
	
	for _, name := range names {
			fmt.Printf("%s\t%d\n", name, ages[name])
	}

	// zero-value of map is nil
	var ages2 map[string]int
	fmt.Println(ages2 == nil)     // true
	fmt.Println(len(ages2) == 0)  // true
	// trying to assign key-value pair to a nil map causes runtime panic
	// ages2["carol"] = 21 // panic: assignment to entry in nil map

	// access a key in a map always yields a value, wether the key exists or not
	// to check if a key is present, one can test it in the following fashion
	// variable ok contains if the key was present or not
	_, ok := ages["alice"]
	if !ok {
		fmt.Println("key 'alice' not found in map")
	}
	// the above can be also written in compact fashion
	// if age, ok := ages["alice"]; !ok { /*...*/ }

	// maps cannot be compared to each other. The only legal comparision is with 'nil'
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))   // false
	// Go does not have a 'set' type. But a map can server as a set using only keys

	// sometimes it may be necessary to use a slice of string as key. But since keys 
	// should be comparable, and slices are not, slices hence cannot be used as map keys
	// a work around can be deviced in two steps
	// 1 - first we make a helper func to map each slice to a string. This can be achieved using Sprintf
	// 2 - whenever accessing the map using  a slice as key, this helper func is called to get the 
	//     slice to string mapping
	var m = make(map[string]int)
	var s := "hello world"
	Add(s[:5], m)
	Add(s[:5], m)
	Add(s[6:], m)
	Add(s[6:], m)
	fmt.Printf("count %s: %d,\t%s : %d\n", s[:5], Count(s[:5]), s[6:], Count(s[:6]))

}

// map a slice to  string and return the string
func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string, m map[string]int) { 
	m[k(list)]++
}

func Count(list []string, m map[string]int) int {
	return m[k(list)]
}

// function to check if to maps are equal
func equal(x, y map[string]int) bool {
	if len(x) != len(y){
		return false
	}
	// iterate of keys and values of map - x 
	for k, xv := range x {
		// if key k is not present in y 
		// or if x[k] and yk] are not same, 
		// return false
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}