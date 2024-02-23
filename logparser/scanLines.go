package logparser

import (
	"log"
	"strconv"
	"strings"
)

func (p LogParser) scanStatusLine(line string, queueItem *QueueItem) {
	matches := p.statusExp.FindStringSubmatch(line)
	if len(matches) == 0 {
		return
	}

	queueItem.Status = matches[1]
	queueItem.StatusLine = line
}

func (p LogParser) scanDSN(line string, queueItem *QueueItem) {
	matches := p.dsnExp.FindStringSubmatch(line)
	if len(matches) == 0 {
		return
	}

	splitter := strings.SplitN(matches[1], ".", 3)

	smtpClass, err := strconv.Atoi(splitter[0])
	if err != nil {
		log.Fatal(err)
	}

	queueItem.SMTPCodeClass = smtpClass

	smtpSubject, err := strconv.Atoi(splitter[1])
	if err != nil {
		log.Fatal(err)
	}

	queueItem.SMTPCodeSubject = smtpSubject

	smtpDetail, err := strconv.Atoi(splitter[2])
	if err != nil {
		log.Fatal(err)
	}

	queueItem.SMTPCodeDetail = smtpDetail
}

func (p LogParser) scanToLine(line string, queueItem *QueueItem) {
	matches := p.toExp.FindStringSubmatch(line)
	if len(matches) == 0 {
		return
	}

	queueItem.To = matches[1]
}
