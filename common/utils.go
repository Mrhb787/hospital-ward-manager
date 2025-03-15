package common

import (
	"fmt"
	"runtime/debug"
)

func Go(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("%v", r)
				fmt.Println("panic err: ", err)
				fmt.Sprintf("recovered from panic: %s", string(debug.Stack()))
				return
			}
		}()
	}()
	f()
}
