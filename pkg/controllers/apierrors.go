package controllers

import "poll/pkg/services"

type ApiError struct {
	Status  int
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}

func ParseError(e error) ApiError {
	switch e {
	case services.ErrorPollNotFound, services.ErrorAnswerNotFound:
		return ApiError{404, e.Error()}
	case services.ErrorNoQuestion, services.ErrorEmptyAnswers:
		return ApiError{400, e.Error()}
	default:
		return ApiError{500, e.Error()}
	}
}
