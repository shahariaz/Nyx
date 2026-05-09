package nyx

import "net/http"

type Nyx struct {
}

func NewNyx() *Nyx {
	return &Nyx{}
}

func (n *Nyx) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from NYX !!"))
}

func (n *Nyx) Run(addr string) error {
	return http.ListenAndServe(addr, n)
}
