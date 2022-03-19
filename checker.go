package fet

import "time"

type Checker interface {
	Check(key string, value interface{}) bool
}

type withCheckerFunc func(key string, value interface{}) bool

func (wcf withCheckerFunc) Check(key string, value interface{}) bool {
	return wcf(key, value)
}

var IfNotNil Checker = withCheckerFunc(func(key string, value interface{}) bool {
	return value != nil
})

var IfNotZero Checker = withCheckerFunc(func(key string, value interface{}) bool {
	return value != 0
})

var IfNotEmpty Checker = withCheckerFunc(func(key string, value interface{}) bool {
	return value != ""
})

var IfNotZeroTime Checker = withCheckerFunc(func(key string, value interface{}) bool {
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
