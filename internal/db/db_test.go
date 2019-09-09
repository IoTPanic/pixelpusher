package db

import (
	"fmt"
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	os.Remove("./tst_db/testing.db")
	os.MkdirAll("./tst_db", 0700) // Create db dir if nonexistant
	err := Connect("./tst_db/testing.db")
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsertAndQueryFixture(t *testing.T) {
	TestConnect(t)
	var err error
	f := Fixture{Name: "tstFixture", LongID: "abcdef123456", PixelsID: 0, UniverseID: 0, ConnectionMethod: "UDP", ConnectionHost: "localhost", ConnectionPort: 1234, ConnectionMQTT: true}
	f.ID, err = f.Insert()
	if err != nil {
		t.Fatal(err)
	}
	tf, err := QueryFixture(f.LongID)
	if err != nil {
		t.Fatal(err)
	}
	if tf != f {
		fmt.Println(tf)
		fmt.Println(f)
		t.Fatal("Query failed, entry not as expected")
	}
}
