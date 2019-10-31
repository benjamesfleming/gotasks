package policies

// Policy interface
type Policy interface {
	CanList() (bool, error)
	CanShow(*interface{}) (bool, error)
	CanCreate() (bool, error)
	CanUpdate(*interface{}) (bool, error)
	CanDestroy(*interface{}) (bool, error)
}
