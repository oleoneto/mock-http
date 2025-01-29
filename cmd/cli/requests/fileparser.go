package requests

import (
	"time"

	"github.com/oleoneto/go-toolkit/helpers"
)

func ParseSchema(debugFunc func(...any), in []byte, parser func([]byte, any) error) (data Schema, err error) {
	debugFunc(helpers.FuncName())

	err = parser(in, &data)
	return
}

type Schema struct {
	Global struct {
		Timeout *time.Duration `yaml:"timeout,omitempty" json:"timeout,omitempty"`
	} `yaml:"global,omitempty" json:"global,omitempty"`
	Requests []Request `yaml:"requests" json:"requests"`
}
