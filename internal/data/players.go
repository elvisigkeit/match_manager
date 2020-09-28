package data

type Position struct {
    X  int `json:"x"`
    Y  int `json:"y"`
}

type Player struct {
    Char 	  string   `json:"char"`
    Pos       Position `json:"pos"`
}

func AddInPlayers(players map[string]Player, id string, char string, x int, y int) {
    player := Player{Char: char, Pos: Position{X: x, Y: y}}
    players[id] = player
}

func ChangePlayerPosition(players map[string]Player, id string, x int, y int) {
    player := players[id]
    player.Pos.X = x
    player.Pos.Y = y
    players[id] = player
}