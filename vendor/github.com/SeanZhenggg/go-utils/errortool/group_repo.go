package errortool

import "log"

type IGroupRepo interface {
	Add(code int)
	Get(code int) int
}

type groupRepo struct {
	store map[int]struct{}
}

func (g *groupRepo) Add(code int) {
	if _, ok := g.store[code]; ok {
		log.Panicf("group error code duplicate definition : %d", code)
	}
	g.store[code] = struct{}{}
}

func (g *groupRepo) Get(code int) int {
	_, ok := g.store[code]

	if !ok {
		log.Panicf("group error code not exist : %d", code)
	}

	return code
}
