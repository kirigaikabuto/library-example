package books

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrNoId = com.NewMiddleError(errors.New("no id in query"), 500, 6)
)
