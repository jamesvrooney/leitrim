package leitrim

import "fmt"

// Where tells me where I'm from
func Where() {
	tullaghan()
}

// ToThePub Informs the user that I'm going to the pub.
func ToThePub() {
	fmt.Println("I'm going to the pub in Tullaghan just down the road.")
}

func tullaghan() {
	fmt.Println("I'm from tullaghan")
}
