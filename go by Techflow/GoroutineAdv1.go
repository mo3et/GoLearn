package main

import (
	"fmt"
)

// channel Select机制

select {
case <- chan1:
case chan2 <-1:
	default
}

func main() {

}
