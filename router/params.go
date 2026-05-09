package router

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (p Params) Get(key string) string {
	for i := range p {
		if p[i].Key == key {
			return p[i].Value
		}
	}
	return ""
}

func (p Params) ByName(name string) string {
	return p.Get(name)
}

func (p Params) Clone() Params {
	if len(p) == 0 {
		return nil
	}
	clone := make(Params, len(p))
	copy(clone, p)
	return clone
}
