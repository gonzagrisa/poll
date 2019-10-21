package services

import (
	"poll/pkg/domain"
	"testing"
)

func TestValidatePoll(t *testing.T) {
	cases := []struct {
		TestName    string
		Poll        domain.Poll
		ExpectedErr error
	}{
		{
			TestName: "Complete Poll",
			Poll: domain.Poll{
				Question: "Test Question",
				Answers:  map[string]string{"1": "test"},
			},
			ExpectedErr: nil,
		},
		{
			TestName: "Without Question",
			Poll: domain.Poll{
				Answers: map[string]string{"1": "test"},
			},
			ExpectedErr: ErrorNoQuestion,
		},
		{
			TestName: "Without Answers",
			Poll: domain.Poll{
				Question: "Test Question",
			},
			ExpectedErr: ErrorEmptyAnswers,
		},
		{
			TestName:    "Empty Poll",
			Poll:        domain.Poll{},
			ExpectedErr: ErrorNoQuestion,
		},
	}

	for _, c := range cases {
		t.Run(c.TestName, func(t *testing.T) {
			result := validatePoll(c.Poll)

			if c.ExpectedErr != result {
				t.Fail()
			}
		})
	}
}

func TestCreateResults(t *testing.T) {
	cases := []struct {
		TestName string
		Poll     domain.Poll
		Expected domain.PollResults
	}{
		{
			TestName: "Complete Poll",
			Poll: domain.Poll{
				Question: "Test Question",
				Answers: map[string]string{
					"1": "test",
					"2": "test 2",
				},
			},
			Expected: domain.PollResults{
				Question: "Test Question",
				Votes: map[string]int{
					"test":   0,
					"test 2": 0,
				},
			},
		},
		{
			TestName: "Without Question",
			Poll: domain.Poll{
				Answers: map[string]string{
					"1": "test",
					"2": "test 2",
				},
			},
			Expected: domain.PollResults{
				Question: "",
				Votes: map[string]int{
					"test":   0,
					"test 2": 0,
				},
			},
		},
		{
			TestName: "Without Answers",
			Poll: domain.Poll{
				Question: "Test Question",
			},
			Expected: domain.PollResults{
				Question: "Test Question",
				Votes:    map[string]int{},
			},
		},
		{
			TestName: "Empty Poll",
			Poll:     domain.Poll{},
			Expected: domain.PollResults{
				Question: "",
				Votes:    map[string]int{},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.TestName, func(t *testing.T) {
			result := createResults(c.Poll)

			if !c.Expected.Equals(result) {
				t.Fail()
			}
		})
	}
}

func TestCreatePoll(t *testing.T) {
	cases := []struct {
		TestName    string
		Poll        domain.Poll
		ExpectedErr error
	}{
		{
			TestName: "Complete Poll",
			Poll: domain.Poll{
				Question: "Test Question",
				Answers: map[string]string{
					"1": "test",
				},
			},
			ExpectedErr: nil,
		},
		{
			TestName: "Without Question",
			Poll: domain.Poll{
				Answers: map[string]string{
					"1": "test",
				},
			},
			ExpectedErr: ErrorNoQuestion,
		},
		{
			TestName: "Without Answers",
			Poll: domain.Poll{
				Question: "Test Question",
			},
			ExpectedErr: ErrorEmptyAnswers,
		},
		{
			TestName:    "Empty Poll",
			Poll:        domain.Poll{},
			ExpectedErr: ErrorNoQuestion,
		},
	}

	for _, c := range cases {
		t.Run(c.TestName, func(t *testing.T) {
			// delete actual poll first
			savedPoll = domain.Poll{}
			savedResults = domain.PollResults{}

			result, err := CreatePoll(c.Poll)
			if c.ExpectedErr != err {
				t.Fail()
			}

			if c.ExpectedErr == nil {
				if !c.Poll.Equals(result) || !c.Poll.Equals(savedPoll) {
					t.Fail()
				}
			}

			// delete created poll
			savedPoll = domain.Poll{}
			savedResults = domain.PollResults{}
		})
	}
}

func TestGetPoll(t *testing.T) {

}
