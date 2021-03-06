package roles

import (
	"github.com/hunghoangmagrabbit/werewolves-of-millers-hollow/game"
)

func InBounds(g *game.Game, i int) bool {
	return i >= 0 && i < g.Count()
}

func Alive(g *game.Game, i int) bool {
	return InBounds(g, i) && g.Player(i).Alive()
}

func NotEqual(g *game.Game, i int, j int) bool {
	return InBounds(g, i) && i != j
}
