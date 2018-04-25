package slackanalyzer

import "time"

// Massage is a slack massage
type Massage struct {
	Sender string
	SendTime time.Time
	Comment string
	Channel string
}
