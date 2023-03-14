package common_wrap

import (
	"github.com/wpliap/common-wrap/config"
	"github.com/wpliap/common-wrap/log"
)

func init() {
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}
	log.InitLog()
}
