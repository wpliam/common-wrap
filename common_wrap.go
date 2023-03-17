package common_wrap

import (
	"github.com/wpliap/common-wrap/config"
)

func init() {
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}
}
