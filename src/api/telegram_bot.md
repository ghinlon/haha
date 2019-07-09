# Telegram Bots 

# Links

* [telebot - GoDoc](https://godoc.org/gopkg.in/tucnak/telebot.v2)
* [tgbotapi - GoDoc](https://godoc.org/github.com/go-telegram-bot-api/telegram-bot-api)

# Create a bot

* [Telegram: Contact @botfather](https://telegram.me/botfather)


# API

* [Telegram Bot API](https://core.telegram.org/bots/api#authorizing-your-bot)
* [tgbotapi - GoDoc](https://godoc.org/github.com/go-telegram-bot-api/telegram-bot-api)
* [Creating a Bot using the Telegram Bot API – Bot Tutorials](https://tutorials.botsfloor.com/creating-a-bot-using-the-telegram-bot-api-5d3caed3266d?gi=7b34239cc74c)

# More

* `getUserProfilePhotos`  
	Use this method to get a list of profile pictures for a user. Returns a `UserProfilePhotos` object.

# Privacy mode

* [Bots: An introduction for developers](https://core.telegram.org/bots#privacy-mode)

# getUpdates

* [Push technology: Long polling - Wikipedia](https://en.wikipedia.org/wiki/Push_technology#Long_polling)

# type Update struct  

```go
type Update struct {
    UpdateID           int                 `json:"update_id"`
    Message            *Message            `json:"message"`
    EditedMessage      *Message            `json:"edited_message"`
    ChannelPost        *Message            `json:"channel_post"`
    EditedChannelPost  *Message            `json:"edited_channel_post"`
    InlineQuery        *InlineQuery        `json:"inline_query"`
    ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
    CallbackQuery      *CallbackQuery      `json:"callback_query"`
    ShippingQuery      *ShippingQuery      `json:"shipping_query"`
    PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
}

type UpdatesChannel <-chan Update
	func (ch UpdatesChannel) Clear()

func (bot *BotAPI) GetUpdates(config UpdateConfig) ([]Update, error)
func (bot *BotAPI) GetUpdatesChan(config UpdateConfig) (UpdatesChannel, error)

type UpdateConfig
    func NewUpdate(offset int) UpdateConfig
```

# func (*BotAPI) Send 

* sendMessage  
	Use this method to send text messages. On success, the sent Message is returned.

```go
func (bot *BotAPI) Send(c Chattable) (Message, error)

// Chattable is any config type that can be sent.
type Chattable interface {
	values() (url.Values, error)
	method() string
}

type Message struct {
    MessageID             int                `json:"message_id"`
    From                  *User              `json:"from"` // optional
    Date                  int                `json:"date"`
    Chat                  *Chat              `json:"chat"`
	... 	// balabala a lot.
}
    func (m *Message) Command() string
    func (m *Message) CommandArguments() string
    func (m *Message) CommandWithAt() string
    func (m *Message) IsCommand() bool
    func (m *Message) Time() time.Time

type MessageConfig
    func NewMessage(chatID int64, text string) MessageConfig
    func NewMessageToChannel(username string, text string) MessageConfig


// values returns a url.Values representation of MessageConfig.
func (config MessageConfig) values() (url.Values, error) {
	v, err := config.BaseChat.values()
	if err != nil {
		return v, err
	}
	v.Add("text", config.Text)
	v.Add("disable_web_page_preview", strconv.FormatBool(config.DisableWebPagePreview))
	if config.ParseMode != "" {
		v.Add("parse_mode", config.ParseMode)
	}

	return v, nil
}

// method returns Telegram API method name for sending Message.
func (config MessageConfig) method() string {
	return "sendMessage"
}
```

**It's confused me a lot of time, `type Message struct` really is not
a `Chattable`, but `MessageConfig` does.**

# Framework

* [telebot - GoDoc](https://godoc.org/gopkg.in/tucnak/telebot.v2)


