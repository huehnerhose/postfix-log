package logparser

type QueueItemCollection map[QueueID]*QueueItem

func (qi QueueItemCollection) String() string {
	var str string
	for _, v := range qi {
		str += v.String()
	}
	return str
}

func (qi QueueItemCollection) FilterByStatusString(status string) QueueItemCollection {
	filtered := make(QueueItemCollection)
	for k, v := range qi {
		if v.Status == status {
			filtered[k] = v
		}
	}
	return filtered
}

func (qi QueueItemCollection) FilterByStatusClass(statusClass int) QueueItemCollection {
	filtered := make(QueueItemCollection)
	for k, v := range qi {
		if v.SMTPCodeClass == statusClass {
			filtered[k] = v
		}
	}
	return filtered
}
