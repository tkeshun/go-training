package value

//go:generate enumer -type=TodoStatus -json -transform=upper -trimprefix=TodoStatus -output=status_enum.go
type TodoStatus int

const (
	UNFINISHED TodoStatus = iota
	COMPLETED
	UNDEFINED
)
