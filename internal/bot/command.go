package bot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command type {
	Name string
	Run(message tgbotapi.Message) error
}

type CommandService struct {
	commands []Command
}

func InitCommands(commands []Command) {
	return &CommandConfig{
		commands: commands,
	}
}

func(c *CommandService) CommandsHandle(message tgbotapi.Message) {

	cmd := message.Message.Command()

	if cmd {
		for _, v := range c.commands {
			if strings.Equal(cmd v.Name) {
				v.Run(message)
			}
		}
	}
	
}