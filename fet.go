// Package fet allows you to build queries in a dynamic, and programmatic way.
//
// You can build mongo queries with function calls,
// instead of writing nested maps. And you have a unique,
// built-in logical control on your queries easily.
package fet

// Updater is an interface that defines how to update a filter.
type Updater interface {
	Update(m M)
}

// Apply applies the given updaters to the given filter.
//
// Example:
// 	filter := fet.Apply(
//		m,
//		fet.Field("firstname").Eq("dave"),
//		fet.Field("lastname").Eq("bowman"),
// 	)
//
func Apply(filter M, updaters ...Updater) M {
	for _, updater := range updaters {
		updater.Update(filter)
	}
	return filter
}

// Build creates a filter and apply the given updaters to it.
//
// Example:
// 	filter := fet.Build(
//		fet.Field("firstname").Eq("dave"),
//		fet.Field("lastname").Eq("bowman"),
// 	)
//
func Build(updaters ...Updater) M {
	return Apply(M{}, updaters...)
}

// BuildSet combines Build and Set. is equivalent to:
// 	fet.Build(fet.Set(updaters...))
func BuildSet(updaters ...Updater) M {
	return Build(Set(updaters...))
}
