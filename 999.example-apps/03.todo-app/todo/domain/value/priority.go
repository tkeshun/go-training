package value

//go:generate enumer -type=TodoPriority -json -transform=upper -trimprefix=TodoPriority -output=priority_enum.go
type TodoPriority int

const (
	LOW TodoPriority = iota
	MEDIUM
	HIGH
)
