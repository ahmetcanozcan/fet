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
	// Set is updater for $set query
	Set = getGlobalOperationWithKeywordConstructor(KeywordSet)
	// SetOnInsert is updater for $setOnInsert query
	SetOnInsert = getGlobalOperationWithKeywordConstructor(KeywordSetOnInsert)
	// Unset is updater for $unset query
	Unset = getGlobalOperationWithKeywordConstructor(KeywordUnset)
)
