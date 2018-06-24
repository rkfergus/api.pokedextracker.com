package games

type IGameService interface {
	List() func() []Game
}

type GameService struct{}
