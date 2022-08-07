package alerts

import (
	"time"

	"anime-skip.com/public-api/internal"
	"github.com/bwmarrin/discordgo"
)

type DiscordAPIClient struct {
	bot       *discordgo.Session
	channelID string
}

func NewDiscordAPIClient(token string, channelID string) (internal.Alerter, error) {
	session, err := discordgo.New("Bot " + token)
	return &DiscordAPIClient{
		bot:       session,
		channelID: channelID,
	}, err
}

func (c *DiscordAPIClient) sendMessage(message string) (*discordgo.Message, error) {
	return c.bot.ChannelMessageSend(c.channelID, message)
}

// CreateThread implements internal.Alerter
func (c *DiscordAPIClient) CreateThread(threadName string, messageContent string) (threadID string, err error) {
	message, err := c.sendMessage(messageContent)
	if err != nil {
		return "", err
	}
	archiveAfter := int((7 * 24 * time.Hour).Minutes())
	thread, err := c.bot.MessageThreadStart(c.channelID, message.ID, threadName, archiveAfter)
	return thread.ID, err
}

// Send implements internal.Alerter
func (c *DiscordAPIClient) Send(message string) error {
	_, err := c.sendMessage(message)
	return err
}

// SendToThread implements internal.Alerter
func (c *DiscordAPIClient) SendToThread(threadID string, message string) error {
	_, err := c.bot.ChannelMessageSend(threadID, message)
	return err
}
