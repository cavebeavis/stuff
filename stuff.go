package stuff

import (
	"fmt"
)

var (
	hey string
)

func init() {
	hey = "hey hey hey"
}

func Stuff() string {
	return fmt.Sprintf("%#v", hey)
}