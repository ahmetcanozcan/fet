package fet

type Updater interface {
	Update(filter M)
}

func Apply(filter M, updaters ...Updater) M {
	for _, updater := range updaters {
		updater.Update(filter)
	}
	return filter
}

func Build(updaters ...Updater) M {
	return Apply(M{}, updaters...)
}

func BuildSet(updaters ...Updater) M {
	return Build(Set(updaters...))
}
