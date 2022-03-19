package fet

type updaterGroup []Updater

func (g updaterGroup) Update(m M) {
	Apply(m, g...)
}

// Group returns a Updater that applies the given updaters.
func Group(u ...Updater) Updater {
	return updaterGroup(u)
}
