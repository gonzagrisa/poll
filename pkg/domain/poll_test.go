package domain

import "testing"

func TestPoll_HasQuestion(t *testing.T) {
	cases := []struct {
		TestName       string
		Poll           Poll
		ExpectedResult bool
	}{
		{
			TestName: "Complete Poll",
			Poll: Poll{
				Question: "Test Question",
				Answers:  map[string]string{
					"1": "test",
				},
			},
			ExpectedResult: true,
		},
		{
			TestName: "Without Answers",
			Poll: Poll{
				Question: "Test Question",
			},
			ExpectedResult: true,
		},
		{
			TestName: "Without Question",
			Poll: Poll{
				Answers: map[string]string{
					"1": "test",
				},
			},
			ExpectedResult: false,
		},
		{
			TestName:       "Empty Poll",
			Poll:           Poll{},
			ExpectedResult: false,
		},
	}

	for _, c := range cases {
		t.Run(c.TestName, func(t *testing.T) {
			result := c.Poll.HasQuestion()

			if c.ExpectedResult != result {
				t.Fail()
			}
		})
	}
}

func TestPoll_HasAnswers(t *testing.T) {
	cases := []struct {
		TestName       string
		Poll           Poll
		ExpectedResult bool
	}{
		{
			TestName: "Complete Poll",
			Poll: Poll{
				Question: "Test Question",
				Answers:  map[string]string{
					"1": "test",
				},
			},
			ExpectedResult: true,
		},
		{
			TestName: "Without Question",
			Poll: Poll{
				Answers: map[string]string{
					"1": "test",
				},
			},
			ExpectedResult: true,
		},
		{
			TestName: "Without Answers",
			Poll: Poll{
				Question: "Test Question",
			},
			ExpectedResult: false,
		},
		{
			TestName:       "Empty Poll",
			Poll:           Poll{},
			ExpectedResult: false,
		},
	}

	for _, c := range cases {
		t.Run(c.TestName, func(t *testing.T) {
			result := c.Poll.HasAnswers()

			if c.ExpectedResult != result {
				t.Fail()
			}
		})
	}
}
