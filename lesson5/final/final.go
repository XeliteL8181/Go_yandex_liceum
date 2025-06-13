package final

import (
	"sort"
	"strings"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func ParseLine(line string) (name, strGoals, strMisses, strAssists string) {
	parts := strings.Fields(line)
	if len(parts) != 4 {
		return "", "", "", ""

	}
	return parts[0], parts[1], parts[2], parts[3]
}

func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
	}

	p.Rating = p.calculateRating()
	return p
}

func (p *Player) calculateRating() float64 {
	base := float64(p.Goals) + float64(p.Assists) / 2
	if p.Misses == 0 {
		return base
	}
	return base / float64(p.Misses)
}

func goalsSort(players []Player) []Player {
	top := append([]Player(nil), players...)
	sort.Slice(top, func(i, j int) bool {
		if top[i].Goals != top[j].Goals {
			return top[i].Goals > top[j].Goals
		}
		return top[i].Name < top[j].Name
	})

	return top
}

func ratingSort(players []Player) []Player {
	top := append([]Player(nil), players...)
	sort.Slice(top, func(i, j int) bool {
		if top[i].Rating != top[j].Rating {
			return top[i].Rating > top[j].Rating
		}
		return top[i].Name < top[j].Name
	})

	return top
}

func gmSort(players []Player) []Player {
	top := append([]Player(nil), players...)

	sort.Slice(top, func(i, j int) bool {
		ratiosI := top[i].Goals
		if top[i].Misses != 0 {
			ratiosI /= top[i].Misses
		}
		ratiosJ := top[j].Goals
		if top[j].Misses != 0 {
			ratiosJ /= top[j].Misses
		}

		if ratiosI != ratiosJ {
			return ratiosI > ratiosJ
		}
		return top[i].Name < top[j].Name
	})

	return top
}
