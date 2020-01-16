package mongodb_sql_driver

import (
	"context"
	"database/sql/driver"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// mConnector implements sql.Conn interface
type mConnector struct {
	mongo *mongo.Client
	cfg   *Config
}

func newMConnector(cfg *Config, mongo *mongo.Client) *mConnector {
	return &mConnector{mongo: mongo, cfg: cfg}
}

func (m *mConnector) Ping(ctx context.Context) error {
	if m.mongo == nil {
		return driver.ErrBadConn
	}
	ctx, _ = context.WithTimeout(ctx, m.cfg.PingTimeout)
	err := m.mongo.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

func (m *mConnector) Prepare(query string) (driver.Stmt, error) {
	return nil, fmt.Errorf("prepare method not implemented")
}

func (m *mConnector) Close() error {
	return nil
}

func (m *mConnector) Begin() (driver.Tx, error) {
	return m, fmt.Errorf("begin method not implemented")
}

func (m *mConnector) Commit() error {
	return fmt.Errorf("commit method not implemented")
}

func (m *mConnector) Rollback() error {
	return fmt.Errorf("rollback method not implemented")
}

func (m *mConnector) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	return nil, fmt.Errorf("begin queryContext not implemented")
}
