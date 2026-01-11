package events

import (
	"dondarrion91/github-activity/internal/utils"
	"fmt"
)

type IPushEvent interface {
	IEvent
}

type PushEvent struct {
	Event
	before  string
	head    string
	commits int
}

func (p PushEvent) String() string {
	return fmt.Sprintf("Pushed %d commits to %s", p.commits, p.repo)
}

func getCommits(repo string, before string, head string) (int, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/compare/%s...%s", repo, before, head)

	commits, err := utils.GetCommits(url)

	if err != nil {
		return 0, err
	}

	return commits, nil
}

func NewPushEvent(repo string, rtype string, payload map[string]interface{}) IPushEvent {
	before := payload["before"].(string)
	head := payload["head"].(string)
	commits, _ := getCommits(repo, before, head)

	return &PushEvent{
		Event: Event{
			repo:  repo,
			rtype: rtype,
		},
		before:  before,
		head:    head,
		commits: commits,
	}
}
