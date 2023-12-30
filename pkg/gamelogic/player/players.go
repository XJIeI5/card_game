package player

type Players []*Player

func (ps *Players) Reset() {
	for _, p := range *ps {
		p.isPassed = false
	}
}
