package slacker

import (
	"context"

	"github.com/slack-go/slack"
)

// A BotContext interface is used to respond to an event
type BotContext interface {
	Context() context.Context
	Event() *slack.MessageEvent
	RTM() *slack.RTM
	Client() *slack.Client
	GetUserName() (string, error)
	GetChannelName() (string, error)
}

// NewBotContext creates a new bot context
func NewBotContext(ctx context.Context, event *slack.MessageEvent, client *slack.Client, rtm *slack.RTM) BotContext {
	return &botContext{ctx: ctx, event: event, client: client, rtm: rtm}
}

type botContext struct {
	ctx    context.Context
	event  *slack.MessageEvent
	client *slack.Client
	rtm    *slack.RTM
}

// Context returns the context
func (r *botContext) Context() context.Context {
	return r.ctx
}

// Event returns the slack message event
func (r *botContext) Event() *slack.MessageEvent {
	return r.event
}

// RTM returns the RTM client
func (r *botContext) RTM() *slack.RTM {
	return r.rtm
}

// Client returns the slack client
func (r *botContext) Client() *slack.Client {
	return r.client
}

func (r *botContext) GetUserName() (string, error) {
	u, err := r.client.GetUserInfo(r.event.User)
	if err != nil {
		return "", err
	}
	return u.Name, nil
}

func (r *botContext) GetChannelName() (string, error) {
	c, err := r.client.GetConversationInfo(r.event.Channel, false)
	if err != nil {
		return "", err
	}
	return c.Name, nil
}