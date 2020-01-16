package mongodb_sql_driver

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	sql.Register("mongodb", new(mDriver))
}

// mDriver implements sql.Driver interface
type mDriver struct {
}

// Open returns new db connection
func (d *mDriver) Open(dsn string) (driver.Conn, error) {
	cfg, err := Parse(dsn)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, cfg.Timeout)
	if len(cfg.MongoDSN) == 0 {
		return nil, fmt.Errorf("MongoDB DSN is empty! Example: &mongoDSN=mongodb+srv://user:password@test.mongodb.net/test")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDSN))
	if err != nil {
		return nil, err
	}
	return newMConnector(cfg, client), nil
}
