package main

import "fmt"

var lg = fmt.Println
var lf = fmt.Printf

func lv(v any) {
	fmt.Printf("%#+v\n\n", v)
}
