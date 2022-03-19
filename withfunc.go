package fet

type WithFunc func(M)
type withFuncConstructor func(key string, value interface{}, checkers ...Checker) Updater

func (w WithFunc) Update(filter M) {
	w(filter)
}

var (
	withValueNotEq = getKeywordWithFuncConstructor(KeywordNe)
	withValueEq    = getKeywordWithFuncConstructor(KeywordEq)

	withValueIn           = getKeywordWithFuncConstructor(KeywordIn)
	withValueNotIn        = getKeywordWithFuncConstructor(KeywordNin)
	withValueElemMatch    = getKeywordWithFuncConstructor(KeywordElemMatch)
	withValueAll          = getKeywordWithFuncConstructor(KeywordAll)
	withValueAllElemMatch = getKeywordWithFuncConstructor(KeywordAllElemMatch)

	withValueSize = getKeywordWithFuncConstructor(KeywordSize)

	withValueGt  = getKeywordWithFuncConstructor(KeywordGt)
	withValueGte = getKeywordWithFuncConstructor(KeywordGte)
	withValueLt  = getKeywordWithFuncConstructor(KeywordLt)
	withValueLte = getKeywordWithFuncConstructor(KeywordLte)

	withValueExists = getKeywordWithFuncConstructor(KeywordExists)
)

func withValueIs(key string, value interface{}, checkers ...Checker) Updater {
	return WithFunc(func(filter M) {
		if checkFilter(key, value, checkers...) {
			filter[key] = value
		}
	})
}

func getKeywordWithFuncConstructor(keyword string, checkers ...Checker) withFuncConstructor {
	return func(key string, value interface{}, checkers ...Checker) Updater {
		return withValueKeyword(key, value, keyword, checkers...)
	}
}

func withValueKeyword(key string, value interface{}, keyword string, checkers ...Checker) Updater {
	return WithFunc(func(filter M) {
		if checkFilter(key, value, checkers...) {
			filter[key] = M{keyword: value}
		}
	})
}

// type check
var _ Updater = WithFunc(nil)
