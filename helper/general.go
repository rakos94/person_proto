package helper

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func CatchError(e *error) {
	if err := recover(); err != nil {
		*e = fmt.Errorf("%v", err)
	}
}

func CatchErrorGeneral() {
	if err := recover(); err != nil {
		log.Println("Error =>", err)
	}
}

// RPCErrDesc get decription error of rpc
func RPCErrDesc(err error) error {
	desc := strings.Split(err.Error(), "desc = ")
	return errors.New(desc[1])
}
