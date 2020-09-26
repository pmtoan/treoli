package trello

import (
	"github.com/adlio/trello"
	"github.com/pmtoan/treoli/config"
	"github.com/pmtoan/treoli/lib"
	"github.com/spf13/viper"
)

func (cmd *Command) GetMe() (*lib.Account, error) {
	token, err := cmd.Client.GetToken(viper.GetString(config.TrelloTokenConfigName), trello.Defaults())
	if err != nil {
		return nil, err
	}
	member, err := cmd.Client.GetMember(token.IDMember, trello.Defaults())
	if err != nil {
		return nil, err
	}

	return &lib.Account{
		Username: member.Username,
		Fullname: member.FullName,
		Email:    member.Email,
	}, nil
}
