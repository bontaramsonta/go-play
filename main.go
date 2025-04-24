package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch n := n.(type) {
	case directMessage:
		return n.senderUsername, n.importance()
	case groupMessage:
		return n.groupName, n.importance()
	case systemAlert:
		return n.alertCode, n.importance()
	default:
		return "", 0
	}
}
