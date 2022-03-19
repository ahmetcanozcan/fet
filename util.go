package fet

type updaterGroup []Updater

func (g updaterGroup) Update(m M) {
	Apply(m, g...)
}

func Group(u ...Updater) Updater {
	return updaterGroup(u)
}
