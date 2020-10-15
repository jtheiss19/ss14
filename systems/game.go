package systems

type GameState int

const (
	GAME_ACTIVE GameState = iota
	GAME_MENU
	GAME_WIN
)

func (gs GameState) String() string {
	return [...]string{"GAME_ACTIVE", "GAME_MENU", "GAME_WIN"}[gs]
}

// Game holds all game-related state and functionality.
// Combines all game-related data into a single class for
// easy access to each of the components and manageability.
type Game struct {
	State         GameState
	keys          []bool
	Width, Height int
}

func NewGame(width, height int) *Game {
	return &Game{Width: width, Height: height, State: GAME_ACTIVE}
}

func (g *Game) ProcessInput(dt float64) {
}

func (g *Game) Update(dt float64) {
}

func (g *Game) Render() {
}

func (g *Game) Init() {
}
