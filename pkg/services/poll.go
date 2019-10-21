package services

import (
	"errors"
	"poll/pkg/domain"
)

var (
	// db
	savedPoll    domain.Poll
	savedResults domain.PollResults

	// errors
	ErrorNoQuestion     = errors.New("ask something please")
	ErrorEmptyAnswers   = errors.New("people need answers")
	ErrorPollNotFound   = errors.New("create a poll first")
	ErrorAnswerNotFound = errors.New("invalid answer")
)

type PollError struct {
	Status  int
	Message string
}

func CreatePoll(poll domain.Poll) (domain.Poll, error) {
	if err := validatePoll(poll); err != nil {
		return poll, err
	}

	savedPoll = poll
	savedResults = createResults(poll)

	return poll, nil
}

func GetPoll() (domain.Poll, error) {
	if err := validatePoll(savedPoll); err != nil {
		return savedPoll, ErrorPollNotFound
	}

	return savedPoll, nil
}

func AnswerPoll(key string) error {
	answer, exists := savedPoll.Answers[key]
	if !exists {
		return ErrorAnswerNotFound
	}

	savedResults.Votes[answer]++

	return nil
}

func GetResults() (domain.PollResults, error) {
	if err := validatePoll(savedPoll); err != nil {
		return savedResults, ErrorPollNotFound
	}

	return savedResults, nil
}

func DeletePoll() {
	savedPoll = domain.Poll{}
	savedResults = domain.PollResults{}
}

func validatePoll(poll domain.Poll) error {
	if !poll.HasQuestion() {
		return ErrorNoQuestion
	}

	if !poll.HasAnswers() {
		return ErrorEmptyAnswers
	}

	return nil
}

func createResults(poll domain.Poll) domain.PollResults {
	results := domain.PollResults{
		Question: poll.Question,
		Votes:    make(map[string]int),
	}

	for _, answer := range poll.Answers {
		results.Votes[answer] = 0
	}

	return results
}
