package slackanalyzer

import "time"

// Message is a slack message
type Message struct {
	Sender string
	SendTime time.Time
	Comment string
	Channel string
}
