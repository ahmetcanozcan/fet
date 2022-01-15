package fet

type WithFunc func(M)
type withFuncConstructor func(key string, value interface{}, checkers ...Checker) Updater

func (w WithFunc) Update(filter M) {
	w(filter)
}

var (
	WithValueNotEq = getKeywordWithFuncConstructor(KeywordNe)
	WithValueEq    = getKeywordWithFuncConstructor(KeywordEq)

	WithValueIn           = getKeywordWithFuncConstructor(KeywordIn)
	WithValueNotIn        = getKeywordWithFuncConstructor(KeywordNin)
	WithValueElemMatch    = getKeywordWithFuncConstructor(KeywordElemMatch)
	WithValueAll          = getKeywordWithFuncConstructor(KeywordAll)
	WithValueAllElemMatch = getKeywordWithFuncConstructor(KeywordAllElemMatch)

	WithValueSize = getKeywordWithFuncConstructor(KeywordSize)

	WithValueGt  = getKeywordWithFuncConstructor(KeywordGt)
	WithValueGte = getKeywordWithFuncConstructor(KeywordGte)
	WithValueLt  = getKeywordWithFuncConstructor(KeywordLt)
	WithValueLte = getKeywordWithFuncConstructor(KeywordLte)

	WithValueExists = getKeywordWithFuncConstructor(KeywordExists)
)

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
