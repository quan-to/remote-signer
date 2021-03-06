package keymagic

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/quan-to/chevron/internal/agent"
	"github.com/quan-to/chevron/internal/config"
	"github.com/quan-to/chevron/internal/tools"
	"github.com/quan-to/chevron/pkg/models"
	"github.com/quan-to/chevron/test"
	"github.com/quan-to/slog"
)

func TestPKSGetKey(t *testing.T) {
	config.PushVariables()
	defer config.PopVariables()
	config.DatabaseDialect = "memory"

	// Test Internal
	z, err := ioutil.ReadFile("../../test/data/testkey_privateTestKey.gpg")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	dbh, err := agent.MakeDatabaseHandler(slog.Scope("TEST"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if dbh == nil {
		t.Fatal("expected database handler to test")
	}

	ctx := context.WithValue(context.Background(), tools.CtxDatabaseHandler, dbh)
	gpgKey, _ := models.AsciiArmored2GPGKey(string(z))

	_, _, err = dbh.AddGPGKey(gpgKey)
	if err != nil {
		t.Errorf("Fail to add key to database: %s", err)
		t.FailNow()
	}

	key, _ := PKSGetKey(ctx, gpgKey.FullFingerprint)

	fp, err := tools.GetFingerPrintFromKey(key)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if !tools.CompareFingerPrint(gpgKey.FullFingerprint, fp) {
		t.Errorf("Expected %s got %s", gpgKey.FullFingerprint, fp)
	}

	// Test External
	config.EnableDatabase = false
	config.SKSServer = "https://keyserver.ubuntu.com/"

	ctx = context.Background()
	key, _ = PKSGetKey(ctx, test.ExternalKeyFingerprint)

	fp, err = tools.GetFingerPrintFromKey(key)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if !tools.CompareFingerPrint(test.ExternalKeyFingerprint, fp) {
		t.Errorf("Expected %s got %s", test.ExternalKeyFingerprint, fp)
	}
}

func TestPKSSearchByName(t *testing.T) {
	config.PushVariables()
	defer config.PopVariables()

	// Test Panics
	config.EnableDatabase = false
	_, err := PKSSearchByName(context.Background(), "", 0, 1)
	if err == nil {
		t.Fatalf("Search should fail as not implemented for rethinkdb disabled!")
	}
}

func TestPKSSearchByFingerPrint(t *testing.T) {
	config.PushVariables()
	defer config.PopVariables()

	// Test Panics
	_, err := PKSSearchByFingerPrint(context.Background(), "", 0, 1)
	if err == nil {
		t.Fatalf("Search should fail as not implemented for rethinkdb disabled!")
	}
}

func TestPKSSearchByEmail(t *testing.T) {
	config.PushVariables()
	defer config.PopVariables()

	// Test Panics
	_, err := PKSSearchByEmail(context.Background(), "", 0, 1)
	if err == nil {
		t.Fatalf("Search should fail as not implemented for rethinkdb disabled!")
	}
}

func TestPKSSearch(t *testing.T) {
	// TODO: Implement method and test
	_, err := PKSSearch(context.Background(), "", 0, 1)
	if err == nil {
		t.Fatalf("Search should fail as not implemented for rethinkdb disabled!")
	}
}

func TestPKSAdd(t *testing.T) {
	config.PushVariables()
	defer config.PopVariables()
	dbh, err := agent.MakeDatabaseHandler(slog.Scope("TEST"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if dbh == nil {
		t.Fatal("expected database handler to test")
	}

	ctx := context.WithValue(context.Background(), tools.CtxDatabaseHandler, dbh)
	// Test Internal
	z, err := ioutil.ReadFile("../../test/data/testkey_privateTestKey.gpg")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fp, err := tools.GetFingerPrintFromKey(string(z))

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	o := PKSAdd(ctx, string(z))

	if o != "OK" {
		t.Errorf("Expected %s got %s", "OK", o)
	}

	p, _ := PKSGetKey(ctx, fp)

	if p == "" {
		t.Errorf("Key was not found")
		t.FailNow()
	}

	fp2, err := tools.GetFingerPrintFromKey(string(p))

	if err != nil {
		t.Error(err)
		t.Error(fmt.Errorf("key data: %s", string(p)))
		t.FailNow()
	}

	if !tools.CompareFingerPrint(fp, fp2) {
		t.Errorf("Fingerprint does not match. Expected %s got %s", fp, fp2)
	}

	// Test External
	// TODO: How to be a good test without stuffying SKS?
}
