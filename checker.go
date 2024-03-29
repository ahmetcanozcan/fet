// Package fet implements a checker system to make a more dynamic query-building process
//
// Any struct that implements Checker interface, can be used as a checker in filter update processes.
// And package fet has some built-in checkers.
package fet

import "time"

// Checker is an interface that defines
// how to check a filter has to be eliminated or not.
type Checker interface {
	Check(key string, value interface{}) bool
}

// CheckerFunc implements Checker interface for functional usage
type CheckerFunc func(key string, value interface{}) bool

// Check pass parameters to the CheckerFunc itself
func (cf CheckerFunc) Check(key string, value interface{}) bool {
	return cf(key, value)
}

// IfNotNil checks if the given value is not nil.
var IfNotNil Checker = CheckerFunc(func(key string, value interface{}) bool {
	return value != nil
})

// IfNotZero checks if the given value is not zero.
var IfNotZero Checker = CheckerFunc(func(key string, value interface{}) bool {
	return value != 0
})

// IfNotEmpty checks if the given value is not a empty empty string.
var IfNotEmpty Checker = CheckerFunc(func(key string, value interface{}) bool {
	return value != ""
})

// IfNotZeroTime checks if the given value is not zero time.
// if value type is not time.Time then it will return true.
var IfNotZeroTime Checker = CheckerFunc(func(key string, value interface{}) bool {
	var (
		t  time.Time
		ok bool
	)

	if t, ok = value.(time.Time); !ok {
		return true
	}

	return !t.IsZero()
})

func checkFilter(key string, value interface{}, checkers ...Checker) bool {

	for _, checker := range checkers {
		if !checker.Check(key, value) {
			return false
		}
	}

	return true
}
