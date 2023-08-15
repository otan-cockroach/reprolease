package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

const connstr = "postgresql://root@localhost:26257/defaultdb?sslmode=disable"

func doTest(t *testing.T) {
	ctx := context.Background()
	time.Sleep(time.Millisecond * 30)
	var err error
	var conn *pgx.Conn
	// github is slow o_o
	for i := 0; i < 5; i++ {
		conn, err = pgx.Connect(ctx, connstr)
		if err == nil {
			break
		}
		t.Logf("crdb not up, restarting %s", err)
		time.Sleep(time.Second)
	}
	require.NoError(t, err)

	for _, stmt := range []string{
		"DROP TABLE IF EXISTS test_table",
		"CREATE TABLE IF NOT EXISTS test_table(id INT PRIMARY KEY)",
		"INSERT INTO test_table VALUES (1), (2)",
		"DROP TABLE IF EXISTS test_table",
	} {
		time.Sleep(time.Millisecond * time.Duration(rand.Int31n(500)))
		_, err = conn.Exec(ctx, stmt)
		require.NoError(t, err)
	}

	require.NoError(t, conn.Close(ctx))

	time.Sleep(time.Millisecond * time.Duration(rand.Int31n(500)))
	conn, err = pgx.Connect(ctx, connstr)
	require.NoError(t, err)
	_, err = conn.Query(ctx, "SELECT * FROM test_table")
	t.Logf("err: %s", err)
	require.Error(t, err)
}

func TestRepro0(t *testing.T) {
	doTest(t)
}
func TestRepro1(t *testing.T) {
	doTest(t)
}
func TestRepro2(t *testing.T) {
	doTest(t)
}
func TestRepro3(t *testing.T) {
	doTest(t)
}
func TestRepro4(t *testing.T) {
	doTest(t)
}
func TestRepro11(t *testing.T) {
	doTest(t)
}
func TestRepro12(t *testing.T) {
	doTest(t)
}
func TestRepro13(t *testing.T) {
	doTest(t)
}
func TestRepro14(t *testing.T) {
	doTest(t)
}
