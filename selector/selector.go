package selector

var selectors = make(map[string]Selector)

type Selector interface {
	Select(opt ...Option) (interface{}, error)
}

func Register(protocol string, s Selector) {
	selectors[protocol] = s
}

func Get(protocol string) Selector {
	return selectors[protocol]
}
