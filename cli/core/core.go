package cli

import (
	"fmt"

	"github.com/gookit/color"
)

r := color.Red
g := color.Green
c := color.Cyan
y := color.Yellow

func Run(argv ...string) {
	fmt.Print("[graphgdb #]")
	r.Print("command line works")
}
