package sqlc

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

const (
	dbSource = "postgresql://root:simplebank@localhost:5433/simplebank?sslmode=disable"
)

func TestMain(m *testing.M) {
	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer pool.Close()
	testQueries = New(pool)
	testDB = pool

	os.Exit(m.Run())
}
