package logparser

import "fmt"

type QueueID string

type QueueItem struct {
	QueueId         QueueID
	Status          string
	To              string
	LogLines        []string
	StatusLine      string
	SMTPCodeClass   int // dsn=class.subject.detail
	SMTPCodeSubject int
	SMTPCodeDetail  int
}

func (li *QueueItem) String() string {
	return fmt.Sprintf("Queue ID: %s\nStatus: %s\nTo: %s\nSMTP Code: %s\nStatus Line: %v\n", li.QueueId, li.Status, li.To, li.SMTPCodeClass, li.StatusLine)
}
