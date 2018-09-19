package main

import (
	"bitbucket.org/twuillemin/doublemetaphone/pkg/doublemetaphone"
	"fmt"
)

func main(){
	primary, secondary := doublemetaphone.DoubleMetaphone("SMITH")
	fmt.Printf("Metaphones for SMITH: first: %v, second: %v\n", primary, secondary)

	primary, secondary = doublemetaphone.DoubleMetaphone("SMYTHE")
	fmt.Printf("Metaphones for SMYTHE: first: %v, second: %v\n", primary, secondary)

	primary, secondary = doublemetaphone.DoubleMetaphone("SCHMIDT")
	fmt.Printf("Metaphones for SCHMIDT: first: %v, second: %v\n", primary, secondary)
}
