package qwirk

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
)

const (
	testDBName = "qwirk_test"
	testDBUser = "testuser"
	testDBPass = "asdl1234kfjoe3ihxoiuv"
)

func TestDB(t *testing.T) *gorm.DB {
	t.Helper()

	// Create the DB. We first drop the existing DB. The complex SQL
	// statement below evicts any connections to that database so we can
	// drop it.
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=qwirk password=qwirk sslmode=disable")
	if err != nil {
		t.Fatal(err.Error())
	}

	db.Exec("SELECT pg_terminate_backend(pg_stat_activity.pid) " +
		"FROM pg_stat_activity " +
		fmt.Sprintf("WHERE pg_stat_activity.datname = '%s' ", testDBName) +
		"AND pid <> pg_backend_pid();")
	db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", testDBName))
	db.Exec(fmt.Sprintf("CREATE DATABASE %s;", testDBName))
	db.Exec(fmt.Sprintf("DROP USER IF EXISTS %s;", testDBUser))
	db.Exec(fmt.Sprintf("CREATE USER %s PASSWORD '%s';", testDBUser, testDBPass))
	db.Close()
	// Now need to grant privs on tables just created. Setting up defaults
	// beforehand gets really messy so this is simpler. Note we need a new
	// connection since the old one was not connected to our tfmodule_test db we
	// just created so the grant is applied to the postgres db instead.
	db, err = gorm.Open("postgres", fmt.Sprintf("host=localhost port=5432 user=qwirk password=qwirk dbname=%s sslmode=disable", testDBName))
	if err != nil {
		t.Fatal(err.Error())
	}

	db.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO %s;", testDBUser))
	db.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO %s;", testDBUser))

	// Start logging all statements, you can see the output of this in the PG
	// instance logs. For typical docker dev setup you can see this by doing:
	//
	//   $ docker logs -f terraformregistry_postgres_1
	//
	// Super useful as it includes all the true binding args not interpolated
	// strings like gorm logging. Also shows transaction boundaries which is
	// good for verification.
	db.Exec(fmt.Sprintf("ALTER ROLE %s SET log_statement = 'all'", testDBUser))
	db.Close()

	db, err = gorm.Open("postgres", fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable", testDBUser, testDBPass, testDBName))
	if err != nil {
		t.Fatal(err.Error())
	}
	db.AutoMigrate(&Game{}, &Player{})
	return db
}
