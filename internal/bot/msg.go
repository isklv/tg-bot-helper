package bot

type BotMessage interface {
	read() string
	write(string) error
}