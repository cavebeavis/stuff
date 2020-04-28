package stuff

import (
	"fmt"
)

var (
	hey string
)

func init() {
	hey = "hey hey hey"
	fmt.Println("init stuff 1")
}

func Stuff() string {
	return fmt.Sprintf("%#v", hey)
}