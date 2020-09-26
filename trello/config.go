package trello

import "github.com/pmtoan/treoli/config"

// save config to local file
func (cmd *Command) SetConfig(key, value string) bool {
	config.SaveConfig(key, value)
	return true
}
