package main

import (
	"fmt"

	"github.com/gbochora/leetcode/design"
)

func main() {
	itr := design.Constructor("abc", 1)
	fmt.Println(itr.Next())    // return "ab"
	fmt.Println(itr.HasNext()) // return True
	fmt.Println(itr.Next())    // return "ac"
	fmt.Println(itr.HasNext()) // return True
	fmt.Println(itr.Next())    // return "bc"
	fmt.Println(itr.HasNext()) // return False
	fmt.Println(itr.Next())    // return "bc"
	fmt.Println(itr.HasNext()) // return False
	fmt.Println(itr.Next())    // return "bc"
	fmt.Println(itr.HasNext()) // return False
	fmt.Println(itr.Next())    // return "bc"
	fmt.Println(itr.HasNext()) // return False
	fmt.Println(itr.Next())    // return "bc"
	fmt.Println(itr.HasNext()) // return False
	fmt.Println(itr.Next())    // return "bc"
	fmt.Println(itr.HasNext()) // return False
	fmt.Println(itr.Next())    // return "bc"
	fmt.Println(itr.HasNext()) // return False
	fmt.Println(itr.Next())    // return "bc"
	fmt.Println(itr.HasNext()) // return False
}

//abc, abd, abe, acd, ace, ade, bcd, bce, bde, cde
