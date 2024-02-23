package logparser

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

type LogParser struct {
	queueIdExp      *regexp.Regexp
	statusExp       *regexp.Regexp
	toExp           *regexp.Regexp
	bounceReasonExp *regexp.Regexp
	dsnExp          *regexp.Regexp

	queueItems QueueItemCollection
}

type Options struct {
	QueueIdExpString      string
	StatusExpString       string
	ToExpString           string
	BounceReasonExpString string
	DSNExpString          string
}

func NewLogParser() *LogParser {
	defaultOptions := &Options{
		QueueIdExpString:      "(?P<item>[[:xdigit:]]{8}):",
		StatusExpString:       "status=(?P<status>[[:alpha:]]+) ",
		ToExpString:           "to=<(?P<status>[^>]+)>",
		BounceReasonExpString: "said: (?P<reason>.+) \\(in reply to",
		DSNExpString:          "dsn=(?P<dsn>[^,]+),",
	}

	return &LogParser{
		queueIdExp:      regexp.MustCompile(defaultOptions.QueueIdExpString),
		statusExp:       regexp.MustCompile(defaultOptions.StatusExpString),
		toExp:           regexp.MustCompile(defaultOptions.ToExpString),
		bounceReasonExp: regexp.MustCompile(defaultOptions.BounceReasonExpString),
		dsnExp:          regexp.MustCompile(defaultOptions.DSNExpString),

		queueItems: make(QueueItemCollection),
	}
}

func (p LogParser) getQueueItemByLine(line string) *QueueItem {
	queueIdMatches := p.queueIdExp.FindStringSubmatch(line)
	if len(queueIdMatches) == 0 {
		return nil
	}

	id := QueueID(queueIdMatches[1])
	if _, ok := p.queueItems[id]; !ok {
		p.queueItems[id] = &QueueItem{
			QueueId: id,
		}
	}

	p.queueItems[id].LogLines = append(p.queueItems[id].LogLines, line)

	return p.queueItems[id]
}

// ReadLog reads a postfix log file and returns a QueueItemCollection
// It appends to the QueueItemCollection if called multiple times
func (p LogParser) ReadLog(path string) (*QueueItemCollection, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		QueueItem := p.getQueueItemByLine(line)
		if QueueItem == nil {
			// not a mail log line
			continue
		}

		p.scanAppendLine(line, QueueItem)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &p.queueItems, nil
}

func (p LogParser) scanAppendLine(line string, queueItem *QueueItem) {
	p.scanStatusLine(line, queueItem)
	p.scanToLine(line, queueItem)
	p.scanDSN(line, queueItem)
}
