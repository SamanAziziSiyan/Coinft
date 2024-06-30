package bot

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	api     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := api.GetUpdatesChan(u)

	return &TelegramBot{
		api:     api,
		updates: updates,
	}, nil
}

func (bot *TelegramBot) Start() {
	ownerChatID := int64(123456789) // Replace with your chat ID
	bot.sendHelpMessage(ownerChatID)

	for update := range bot.updates {
		if update.Message != nil {
			bot.handleMessage(update.Message)
		} else if update.CallbackQuery != nil {
			bot.handleCallbackQuery(update.CallbackQuery)
		}
	}
}

func (bot *TelegramBot) handleMessage(message *tgbotapi.Message) {
	text := message.Text
	if strings.HasPrefix(text, "/") {
		response := bot.processCommand(message.Chat.ID, text)

		if response != "" {
			msg := tgbotapi.NewMessage(message.Chat.ID, response)
			bot.api.Send(msg)
		}
	}
}

func (bot *TelegramBot) handleCallbackQuery(callbackQuery *tgbotapi.CallbackQuery) {
	data := callbackQuery.Data
	chatID := callbackQuery.Message.Chat.ID

	response := bot.processCommand(chatID, data)
	if response != "" {
		msg := tgbotapi.NewMessage(chatID, response)
		bot.api.Send(msg)
	}

	callback := tgbotapi.NewCallback(callbackQuery.ID, "")
	bot.api.Request(callback)
}

func (bot *TelegramBot) processCommand(chatID int64, command string) string {
	parts := strings.SplitN(command, " ", 2)
	mainCommand := strings.ToLower(parts[0])

	switch mainCommand {
	case "/play":
		bot.sendWebApp(chatID, "Click to Open the WebApp", "https://t.me/TDDBOT_bot/myapp")
		return ""
	case "/about":
		return "This bot is a demo for handling commands and WebApps in Telegram."
	case "/help":
		bot.sendHelpMessage(chatID)
		return ""
	default:
		return "I don't understand that command. Type /help for a list of commands."
	}
}

func (bot *TelegramBot) sendWebApp(chatID int64, text string, url string) {
	webAppButton := tgbotapi.NewInlineKeyboardButtonURL(text, url)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			webAppButton,
		),
	)
	_, err := bot.api.Send(msg)
	if err != nil {
		fmt.Println("Failed to send Web App button:", err)
	}
}

func (bot *TelegramBot) sendHelpMessage(chatID int64) {
	helpText := `Here are the commands you can use:

/play - Open a WebApp with the specified title
/about - Get information about this bot
/help - Show this help message`

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Play", "/play"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("About", "/about"),
			tgbotapi.NewInlineKeyboardButtonData("Help", "/help"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, helpText)
	msg.ReplyMarkup = keyboard

	bot.api.Send(msg)
}

func (bot *TelegramBot) getHelpText() string {
	return `Here are the commands you can use:
    
/play - Open a WebApp with the specified title
/about - Get information about this bot
/help - Show this help message`
}
