// Package fet implements Updaters for non field-related logical keywords
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
	// Or is updater for $or query
	Or = getLogicalOperationWithKeywordConstructor(KeywordOr)
	// And is updater for $and query
	And = getLogicalOperationWithKeywordConstructor(KeywordAnd)
	// Nor is updater for $nor query
	Nor = getLogicalOperationWithKeywordConstructor(KeywordNor)
	// Nand is updater for $nand query
	Nand = getLogicalOperationWithKeywordConstructor(KeywordNand)
)
