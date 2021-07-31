package notifier

import (
	"github.com/signedsecurity/signotifi3r/internal/configuration"
	"github.com/signedsecurity/signotifi3r/pkg/ansi"
	"github.com/valyala/fasthttp"
)

// Notify handles the notification engine
type Notify struct {
	options *configuration.Options
	client  *fasthttp.Client
	slack   *Slack
}

// New notify instance
func New(options *configuration.Options) (*Notify, error) {
	client := &fasthttp.Client{}
	return &Notify{
		options: options,
		client:  client,
		slack: &Slack{
			client:     client,
			webHookURL: options.YAMLConfig.Platforms.Slack.WebHookURL,
		},
	}, nil
}

// SendNotification to registered webhooks
func (n *Notify) SendNotification(message string) error {
	message = ansi.Strip(message)

	if n.options.YAMLConfig.Platforms.Slack.Enabled {
		err := n.slack.send(message)
		if err != nil {
			return err
		}
	}

	return nil
}
