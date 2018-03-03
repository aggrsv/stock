package sql

import (
	"log"
	"testing"
	"time"

	"github.com/pborman/uuid"
)

// TODO: fix test
var (
	dsn = `root:1234@tcp(192.168.1.253:3307)/sql_test?strict=true&sql_notes=false&parseTime=true`
)

const testExecSql = `
INSERT INTO
	sql_test_tb
(
	Name, Secret, Email, Authorize, Status
)
VALUES
(
	?, ?, ?, ?, ?
)
`

func TestExec(t *testing.T) {
	db, err := Open(`mysql`, dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := db.Exec(
		testExecSql,
		1,
		"Matthew",
		"1234",
		"matthewj.he@gmail.com",
		"45483464",
		0,
	); err != nil {
		t.Fatal(err)
	}
}

func TestConcurrentExec(t *testing.T) {
	db, err := Open(`mysql`, dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	n := 200
	c := 10
	nChan := make(chan time.Duration, n)
	cChan := make(chan int, c)
	for i := 0; i < n; i++ {
		cChan <- 1
		go func() {
			now := time.Now()
			defer func() {
				nChan <- time.Now().Sub(now)
				<-cChan
			}()

			if err := db.Exec(
				testExecSql,
				1,
				uuid.New(),
				uuid.New(),
				uuid.New(),
				uuid.New(),
				0,
			); err != nil {
				t.Fatal(err)
			}
		}()
	}

	var max time.Duration
	var total time.Duration
	for i := 0; i < n; i++ {
		rs := <-nChan
		if rs > max {
			max = rs
		}
		total += rs
	}
	log.Println("Max exec time:", max)
	log.Println("Average exec time:", total/time.Duration(n))
}
