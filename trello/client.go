package trello

import "github.com/pmtoan/treoli/lib"

type Client interface {
	SetConfig(key, value string) bool

	// Personal Info
	GetMe() (*lib.Account, error)

	// Board Info
	GetMyBoards() ([]lib.Board, error)
}
