// demonstration of composite type value for map

// graphs are maps whose value is a map of the key-type
// i,e of nodes of graphs are ints, graphs are maps of where
// key (vertex) is the of type int, and value is a map of [int]bool
// the bool usually represent connection. Similarly, 
// if vertices are denoted by strings, graph is a map
// whose keys are of type string and  values are a map of type [string]bool

package main
import "fmt"

/*
A----B
|\   |
|  \ |
C--- D
*/
func main(){
    // nested maps
    var graph = make(map[string]map[string]bool)
    addEdge("A", "B", graph)
    addEdge("A", "C", graph)
    addEdge("A", "D", graph)
    addEdge("B", "D", graph)
    addEdge("C", "D", graph)
    
    for k, v := range graph{
        fmt.Printf("%v: %v\n", k, v)
    }

    fmt.Printf("C - B edge: %v\n", hasEdge("C", "B", graph)) // C - B edge: false
}

func addEdge(from, to string, G map[string]map[string]bool){
    edges := G[from]
    if edges == nil {
        edges = make(map[string]bool)
        G[from] = edges
    }
    edges[to] = true
}

func hasEdge(from, to string, G map[string]map[string]bool) bool {
    return G[from][to]
}
