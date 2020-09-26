package trello

import (
	"github.com/adlio/trello"
	"github.com/pmtoan/treoli/lib"
)

func (cmd *Command) GetMyBoards() ([]lib.Board, error) {
	rawBoards, err := cmd.Client.GetMyBoards(trello.Defaults())
	if err != nil {
		return nil, err
	}

	var boards []lib.Board
	for _, board := range rawBoards {
		boards = append(boards, lib.Board{
			ID:      board.ID,
			Name:    board.Name,
			Desc:    board.Desc,
			Closed:  board.Closed,
			Pinned:  board.Pinned,
			Starred: board.Starred,
			URL:     board.ShortURL,
		})
	}
	return boards, nil
}
