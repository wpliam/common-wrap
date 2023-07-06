package selector

import "github.com/wpliam/common-wrap/registry"

var selectors = make(map[string]Selector)

type Selector interface {
	Select(opt ...Option) (registry.Proxy, error)
}

func Register(protocol string, s Selector) {
	selectors[protocol] = s
}

func Get(protocol string) Selector {
	return selectors[protocol]
}
