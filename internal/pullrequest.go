package events

import "fmt"

type IPullRequestEvent interface {
	IEvent
}

type PullRequestEvent struct {
	Event
	number   int64
	action   string
	assignee string
}

func (p PullRequestEvent) String() string {
	if p.action == "opened" {
		return fmt.Sprintf("Opened new merge request with number %d in %s", p.number, p.repo)
	}

	if p.action == "assigned" {
		return fmt.Sprintf("Assigned merge request with number %d to %s in %s", p.number, p.assignee, p.repo)
	}

	return fmt.Sprintf("Action %s not implemented yet", p.action)
}

func NewPullRequestEvent(repo string, rtype string, payload map[string]interface{}) IPullRequestEvent {
	action := payload["action"].(string)
	// assignee := payload["assignee"]
	number := payload["number"].(float64)
	var assigneeName string

	return &PullRequestEvent{
		Event: Event{
			repo:  repo,
			rtype: rtype,
		},
		action:   action,
		number:   int64(number),
		assignee: assigneeName,
	}
}
