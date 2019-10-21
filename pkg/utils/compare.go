package utils

import "fmt"

func CompareMapsStringString(m1, m2 map[string]string) (equals bool) {
	if m1 == nil && m2 == nil {
		return true
	}

	if m1 == nil || m2 == nil {
		println(fmt.Sprintf("only one nil map: m1 %+v - m2 %+v", m1, m2))
		return false
	}

	if len(m1) != len(m2) {
		println("invalid lens")
		return false
	}

	for k, v := range m1 {
		value, exists := m2[k]

		if !exists || v != value {
			println(fmt.Sprintf("key %s, exists: %t, value1: %s, value2: %s - m1: %+v - m2: %+v", k, exists, v, value, m1, m2))
			return false
		}
	}

	return true
}

func CompareMapsStringInt(m1, m2 map[string]int) (equals bool) {
	if m1 == nil && m2 == nil {
		return true
	}

	if m1 == nil || m2 == nil {
		println(fmt.Sprintf("only one nil map: m1 %+v - m2 %+v", m1, m2))
		return false
	}

	if len(m1) != len(m2) {
		println("invalid lens")
		return false
	}

	for k, v := range m1 {
		value, exists := m2[k]

		if !exists || v != value {
			println(fmt.Sprintf("key %s, exists: %t, value1: %d, value2: %d - m1: %+v - m2: %+v", k, exists, v, value, m1, m2))
			return false
		}
	}

	return true
}
