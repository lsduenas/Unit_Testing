package hunt

import (
	"fmt"
)

type Shark struct {
	hungry bool
	tired  bool
	speed  int
}

type Prey struct {
	name  string
	speed int
}

func (s *Shark) Hunt(p *Prey) (err error) {
	if p == nil {
		err = fmt.Errorf("nil prey")
		return 
	}
	if s.tired {
		err = fmt.Errorf("cannot hunt, i am really tired")
		return 
	}
	if !s.hungry {
		err = fmt.Errorf("cannot hunt, i am not hungry")
		return 
	}
	if p.speed >= s.speed {
		s.tired = true
		err = fmt.Errorf("could not catch it")
		return 
	}

	s.hungry = false
	s.tired = true
	return
}
