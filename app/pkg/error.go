package pkg

import (
	"errors"
	"fmt"
	"github.com/andifg/artemis_backend/app/constant"
)

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constant.ResponseStatus) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}
