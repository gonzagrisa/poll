package utils

import "testing"

func TestCompareMapsStringString(t *testing.T) {
	cases := []struct {
		TestName string
		Map1     map[string]string
		Map2     map[string]string
		Expected bool
	}{
		{
			TestName: "Nil maps",
			Map1:     nil,
			Map2:     nil,
			Expected: true,
		},
		{
			TestName: "Nil m1",
			Map1:     nil,
			Map2: map[string]string{
				"1": "test",
			},
			Expected: false,
		},
		{
			TestName: "Nil m2",
			Map1: map[string]string{
				"1": "test",
			},
			Map2:     nil,
			Expected: false,
		},
		{
			TestName: "Different lens",
			Map1: map[string]string{
				"1": "test",
			},
			Map2: map[string]string{
				"1": "test",
				"2": "test 2",
			},
			Expected: false,
		},
		{
			TestName: "Different keys",
			Map1: map[string]string{
				"1": "test",
				"2": "test 2",
			},
			Map2: map[string]string{
				"1": "test",
				"3": "test 2",
			},
			Expected: false,
		},
		{
			TestName: "Different values",
			Map1: map[string]string{
				"1": "test",
				"2": "test 2",
			},
			Map2: map[string]string{
				"1": "test",
				"2": "test 3",
			},
			Expected: false,
		},
		{
			TestName: "Equals",
			Map1: map[string]string{
				"1": "test",
				"2": "test 2",
			},
			Map2: map[string]string{
				"1": "test",
				"2": "test 2",
			},
			Expected: true,
		},
	}

	for _, c := range cases {
		t.Run(c.TestName, func(t *testing.T) {
			result := CompareMapsStringString(c.Map1, c.Map2)

			if c.Expected != result {
				t.Fail()
			}
		})
	}
}

func TestCompareMapsStringInt(t *testing.T) {
	cases := []struct {
		TestName string
		Map1     map[string]int
		Map2     map[string]int
		Expected bool
	}{
		{
			TestName: "Nil maps",
			Map1:     nil,
			Map2:     nil,
			Expected: true,
		},
		{
			TestName: "Nil m1",
			Map1:     nil,
			Map2: map[string]int{
				"test": 1,
			},
			Expected: false,
		},
		{
			TestName: "Nil m2",
			Map1: map[string]int{
				"test": 1,
			},
			Map2:     nil,
			Expected: false,
		},
		{
			TestName: "Different lens",
			Map1: map[string]int{
				"test": 1,
			},
			Map2: map[string]int{
				"test":   1,
				"test 2": 1,
			},
			Expected: false,
		},
		{
			TestName: "Different keys",
			Map1: map[string]int{
				"test":   1,
				"test 2": 1,
			},
			Map2: map[string]int{
				"test":   1,
				"test 3": 1,
			},
			Expected: false,
		},
		{
			TestName: "Different values",
			Map1: map[string]int{
				"test":   1,
				"test 2": 1,
			},
			Map2: map[string]int{
				"test":   1,
				"test 2": 2,
			},
			Expected: false,
		},
		{
			TestName: "Equals",
			Map1: map[string]int{
				"test":   1,
				"test 2": 2,
			},
			Map2: map[string]int{
				"test":   1,
				"test 2": 2,
			},
			Expected: true,
		},
	}

	for _, c := range cases {
		t.Run(c.TestName, func(t *testing.T) {
			result := CompareMapsStringInt(c.Map1, c.Map2)

			if c.Expected != result {
				t.Fail()
			}
		})
	}
}
