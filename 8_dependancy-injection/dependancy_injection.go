package dependancy_injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) { // fmt.Fprintf allows you to pass in an io.Writer which we know both os.Stdout and bytes.Buffer implement.
	fmt.Fprintf(writer, "Hello, %s", name)
}
