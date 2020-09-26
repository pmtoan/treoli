package trello

import (
	"github.com/adlio/trello"
	"github.com/pmtoan/treoli/config"
	"github.com/spf13/viper"
)

type Command struct {
	Client *trello.Client
}

func InitClient() *Command {
	appKey := viper.GetString(config.TrelloAppKeyConfigName)
	token := viper.GetString(config.TrelloTokenConfigName)
	if appKey == "" || token == "" {
		return nil
	}

	return &Command{Client: trello.NewClient(appKey, token)}
}
