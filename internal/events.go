package events

import "fmt"

type IEvent interface {
	getRepo() string
	setRepo(name string)
	getType() string
	setType(rtype string)
}

type Event struct {
	repo  string
	rtype string
}

func (e *Event) setRepo(name string) {
	e.repo = name
}

func (e *Event) getRepo() string {
	return e.repo
}

func (e *Event) setType(rtype string) {
	e.rtype = rtype
}

func (e *Event) getType() string {
	return e.rtype
}

func GetEvent(action map[string]interface{}) (IEvent, error) {
	repo := action["repo"].(map[string]interface{})
	repoName := repo["name"].(string)
	rtype := action["type"].(string)
	payload := action["payload"].(map[string]interface{})

	if rtype == "PushEvent" {
		return NewPushEvent(repoName, rtype, payload), nil
	}

	if rtype == "PullRequestEvent" {
		return NewPullRequestEvent(repoName, rtype, payload), nil
	}

	return nil, fmt.Errorf("%s type not implemented", rtype)
}
