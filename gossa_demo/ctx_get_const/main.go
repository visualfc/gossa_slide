package main

import (
	"fmt"
	"go/token"
	"log"

	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg/fmt"
	_ "github.com/goplus/gossa/pkg/math"
)

var source = `
package bar

import "math"

const Pi = math.Pi

func add(i, j int) int {
	return i+j
}
`

func main() {
	fset := token.NewFileSet()
	ctx := gossa.NewContext(0)
	pkg, err := ctx.LoadFile(fset, "main.go", source)
	if err != nil {
		log.Panicln("load", err)
	}
	interp, err := ctx.NewInterp(pkg)
	if err != nil {
		log.Panicln("interp", err)
	}

	// go/constant.Value
	if v, ok := interp.GetConst("Pi"); ok {
		fmt.Printf("%v %T\n", v, v)
		fmt.Println(v.ExactString())
	}
}
