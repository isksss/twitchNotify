package commands

import (
	"github.com/bwmarrin/discordgo"
)

var pingCommand = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Replies with Pong!",
}

// All に今後追加するすべてのコマンドを列挙
var All = []*discordgo.ApplicationCommand{
	pingCommand,
}

// Interaction ハンドラー
func HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	switch i.ApplicationCommandData().Name {
	case "ping":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	}
}
