package qwirk

const (
	Ready State = iota
	InProgress
	Done
)

type State int

func (s State) String() string {
	return [...]string{"Ready", "InProgress", "Done"}[s]
}
