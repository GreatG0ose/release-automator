package ms_teams

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendMessage(t *testing.T) {
	err := SendMessage(
		"https://heidelpay.webhook.office.com/webhookb2/b574ac44-cfa5-4cd2-8b23-7c3d52d0acea@797b2bda-888b-44ab-9967-3c9448c99377/IncomingWebhook/c1d68196df5e42f0a28edbfd85245843/f43437ec-7ce7-4ae5-8082-942c9829d10d",
		Message{
			Title:   "Test title",
			Content: "test content with mention",
			Mentions: []Mention{
				{
					Name:    "Vlad",
					TeamsID: "vladislav.lipianin@unzer.com",
				},
			},
		},
	)
	assert.NoError(t, err)
}
