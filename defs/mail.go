package defs

type MailMessage struct {
	From       string
	Msg        []byte
	Recipients []string
}
