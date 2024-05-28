package msdb

type ITransaction interface {
	Action(func(conn IConnection) error) error
}
