package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestDDNS(t *testing.T) {

	DDNS.Call(tgbotapi.Update{
		UpdateID:           0,
		Message:            nil,
		EditedMessage:      nil,
		ChannelPost:        nil,
		EditedChannelPost:  nil,
		InlineQuery:        nil,
		ChosenInlineResult: nil,
		CallbackQuery:      nil,
		ShippingQuery:      nil,
		PreCheckoutQuery:   nil,
	})
}
