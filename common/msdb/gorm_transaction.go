package msdb

import "gorm.io/gorm"

type GormTransaction struct {
	conn IConnection
}

func NewTransaction(db *gorm.DB) ITransaction {
	return &GormTransaction{conn: NewGormConn(db)}
}

func (t *GormTransaction) Action(f func(conn IConnection) error) (err error) {
	t.conn.Begin() // 开启事务

	err = f(t.conn) // 执行

	if err != nil {
		t.conn.Rollback() // 执行过程出现异常，执行回滚操作
	}

	t.conn.Commit() // 提交事务

	return
}
