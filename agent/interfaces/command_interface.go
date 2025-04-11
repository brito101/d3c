package interfaces

type Command interface {
	Exec() (response string, err error)
}
