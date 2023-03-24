package bot

import (
	"Serega_discord_bot/config"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sashabaranov/go-openai"
)

var ID string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	goBot.Identify.Intents = discordgo.IntentsGuildMessages

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	ID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == ID {
		return
	}

	if m.ChannelID == "1088724952745787482" {
		botMessage, err := s.ChannelMessageSend(m.ChannelID, "Жди нахуй, я думаю...")
		messageID := botMessage.ID
		client := openai.NewClient(config.APIToken)
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: m.Content,
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}
		s.ChannelMessageEdit(m.ChannelID, messageID, resp.Choices[0].Message.Content)
	}
}
