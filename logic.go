package fet

func logicalOperatioWithKeyword(keyword string, updaters ...Updater) Updater {
	return WithFunc(func(filter M) {
		var conds []M
		for _, updater := range updaters {
			m := M{}
			updater.Update(m)
			conds = append(conds, m)
		}
		filter[keyword] = conds

	})
}

func getLogicalOperationWithKeywordConstructor(keyword string) func(updaters ...Updater) Updater {
	return func(updaters ...Updater) Updater {
		return logicalOperatioWithKeyword(keyword, updaters...)
	}
}

var (
	Or   = getLogicalOperationWithKeywordConstructor(KeywordOr)
	And  = getLogicalOperationWithKeywordConstructor(KeywordAnd)
	Nor  = getLogicalOperationWithKeywordConstructor(KeywordNor)
	Nand = getLogicalOperationWithKeywordConstructor(KeywordNand)
)
