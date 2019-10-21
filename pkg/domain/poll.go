package domain

import "poll/pkg/utils"

type Poll struct {
	Question string
	Answers  map[string]string
}

func (p Poll) HasQuestion() bool {
	return p.Question != ""
}

func (p Poll) HasAnswers() bool {
	return len(p.Answers) > 0
}

func (p Poll) Equals(p2 Poll) bool {
	return p.Question == p2.Question && utils.CompareMapsStringString(p.Answers, p2.Answers)
}

type PollResults struct {
	Question string
	Votes    map[string]int
}

func (p PollResults) Equals(p2 PollResults) bool {
	return p.Question == p2.Question && utils.CompareMapsStringInt(p.Votes, p2.Votes)
}
