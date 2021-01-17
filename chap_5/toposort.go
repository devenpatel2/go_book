// Topological sorting of string
// example of Depth-Frist Search (DFS) algorithm
// The prere quisites are given in the prereqs table below, which is a mapping from
// each course to the list of courses that must be completed before it.
// This makes an acyclic graph.

package main
import (
    "fmt"
    "sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string {
    "algorithms": {"data structures"},
    "calculus": {"linear algebra"},
    "compilers": {
        "data structures",
        "formal languages",
        "computer organization",
    },

    "data structures": {"discrete math"},
    "databases": {"data structures"},
    "discrete math": {"intro to programming"},
    "formal languages": {"discrete math"},
    "networks": {"operating systems"},
    "operating systems": {"data structures", "computer organization"},
    "programming languages": {"data structures", "computer organization"},
}


func main(){
    for i, course := range topoSort(prereqs) {
        fmt.Printf("%d:\t%s\n", i+1, course)
    }
}

func topoSort(m map[string][]string) []string {
    var order []string
    seen := make(map[string]bool)
    var visitAll  func(items []string)



    visitAll = func(items []string) {
        for _, item := range items{
            if !seen[item]{
                seen[item] = true
                visitAll(m[item])
                order = append(order, item)
            }
        }
    }
    var keys []string
    for key := range m {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    // recursion with anonymous fuction
    // advatange of using anonymous function is that it has access to variable outside it's scope
    // to use recursion with anonymous functions, it has to be called 'outside' the body of function definitiion 
    visitAll(keys)
    return order
}
