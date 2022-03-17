package fet

func globalOperationWithKeyword(keyword string, updaters ...Updater) Updater {
	return WithFunc(func(filter M) {
		m := Build(updaters...)
		Field(keyword).Is(m).Update(filter)
	})
}

func getGlobalOperationWithKeywordConstructor(keyword string) func(updaters ...Updater) Updater {
	return func(updaters ...Updater) Updater {
		return globalOperationWithKeyword(keyword, updaters...)
	}
}

var (
	Set         = getGlobalOperationWithKeywordConstructor(KeywordSet)
	SetOnInsert = getGlobalOperationWithKeywordConstructor(KeywordSetOnInsert)
	Unset       = getGlobalOperationWithKeywordConstructor(KeywordUnset)
)
