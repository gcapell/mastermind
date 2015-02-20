package main

import (
	"fmt"
	"math"
)

const (
	colors = 6
	pegs   = 4
)

type (
	color int
	guess int
	score struct{ black, white int }
)

var (
	allScores []score
	maxGuess  guess
)

func parseGuess(s string) guess {
	var g guess
	for _, r := range s {
		g = g*colors + guess(r-'0')
	}
	return g
}

func (g guess) String() string {
	var reply string
	var c color
	for p := 0; p < pegs; p++ {
		c, g = color(g%colors), g/colors
		reply = string('0'+c) + reply
	}
	return reply
}

func (s score) String() string {
	return fmt.Sprintf("%d,%d", s.black, s.white)
}

func init() {
	for b := 0; b <= pegs; b++ {
		for w := 0; b+w <= pegs; w++ {
			if b+w == pegs && w == 1 {
				continue
			}
			allScores = append(allScores, score{b, w})
		}
	}
	maxGuess = guess(math.Pow(colors, pegs))
}

func (s score) done() bool {
	return s.black == pegs
}

func firstGuess() guess {
	var g guess
	for p := 0; p < pegs/2; p++ {
		g = g*colors + 1
	}
	return g
}

func solve(hidden guess) {
	g := firstGuess()
	eliminated := make(map[guess]bool)
	for {
		s := hidden.judge(g)
		if s.done() {
			break
		}
		elimThisRound := 0
		for p := guess(0); p < maxGuess; p++ {
			if eliminated[p] {
				continue
			}
			if p.judge(g) != s {
				eliminated[p] = true
				elimThisRound++
			}
		}
		fmt.Printf("guess %s, hidden %s, score %s, eliminated %d\n", g, hidden, s, elimThisRound)
		g = nextGuess(eliminated)
	}
}

type key struct{ h, g guess }

var judgeCache = make(map[key]score)

func (hidden guess) judge(g guess) score {
	k := key{hidden, g}
	if s, ok := judgeCache[k]; ok {
		return s
	}
	var s score
	avail, want := make(map[color]int), make(map[color]int)
	for p := 0; p < pegs; p++ {
		var hc, gc color
		hc, hidden = color(hidden%colors), hidden/colors
		gc, g = color(g%colors), g/colors
		if hc == gc {
			s.black++
		} else {
			avail[hc]++
			want[gc]++
		}
	}
	for c, n := range want {
		s.white += min(n, avail[c])
	}
	judgeCache[k] = s
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func nextGuess(eliminated map[guess]bool) guess {
	maxMin := 0
	judged := 0
	var best guess
	for g := guess(0); g < maxGuess; g++ {
		minElim := int(maxGuess)
		for _, s := range allScores {
			elim := 0
			for p := guess(0); p < maxGuess; p++ {
				if eliminated[p] {
					continue
				}
				judged++
				if p.judge(g) != s {
					elim++
				}
			}
			if elim < minElim {
				minElim = elim
			}
		}
		if minElim < maxMin {
			continue
		}
		if minElim == maxMin {
			if _, ok := eliminated[g]; ok {
				continue
			}
		}
		best = g
		maxMin = minElim
	}
	fmt.Printf("guess %s eliminates at least %d/%d (judged %d)\n", 
		best, maxMin, (int(maxGuess) - len(eliminated)), judged)
	return best
}

func main() {
	solve(parseGuess("4321"))
}
