// Demonstration of a defer call
// a defer call can be used to pair "on entry" and "on exit" actions
package main
import (
	"log"
	"time"
)

func bigSlowOperation() {
	// the trace function is called here due to the parenthesis at the end
	// and the returned function is called at the exit point of the current function

	// the defered function is called "after" the return statements have updated 
	// the function's variables
	defer trace("bigSlowOperation")()	// don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second)		// simulate slow operation by sleeping 		
	log.Printf("Done Sleeping, exiting function")
}

// the returned function is called at the exit point of the calling function
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }	
}

func main(){
	bigSlowOperation()
}