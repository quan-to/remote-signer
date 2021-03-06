package rql

import (
	"fmt"
	"time"

	"github.com/quan-to/chevron/pkg/models"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var userTokenTableInit = tableInitStruct{
	TableName:    "tokens",
	TableIndexes: []string{"Token", "Username", "Fingerprint", "CreatedAt"},
}

func (h *RethinkDBDriver) initUserTokenTable() error {
	return h.initFromStruct(userModelTableInit)
}

func (h *RethinkDBDriver) migrateUserTokenTable() error {
	// Legacy fields
	h.log.Info("[UserToken] Migrating old fields to new fields")
	result, err := r.Table(userTokenTableInit.TableName).
		Filter(r.Row.HasFields("Fingerprint").Not()).
		Update(map[string]interface{}{
			"Fingerprint": r.Row.Field("FingerPrint"),
		}).RunWrite(h.conn)

	if err != nil {
		return err
	}

	h.log.Info("[UserToken] Migrated %d tokens FingerPrint -> Fingerprint", result.Updated)
	return nil
}

// AddUserToken adds a new user token to be valid and returns its token ID
func (h *RethinkDBDriver) AddUserToken(ut models.UserToken) (string, error) {
	rut, err := convertToRethinkDB(ut)
	if err != nil {
		return "", err
	}
	wr, err := r.Table(userTokenTableInit.TableName).
		Insert(rut).
		RunWrite(h.conn)

	if err != nil {
		return "", err
	}

	return wr.GeneratedKeys[0], err
}

// RemoveUserToken removes a user token from the database
func (h *RethinkDBDriver) RemoveUserToken(token string) (err error) {
	_, err = r.Table(userTokenTableInit.TableName).
		GetAllByIndex("Token", token).
		Limit(1).
		Delete().
		RunWrite(h.conn)

	if err != nil {
		return err
	}

	return nil
}

// GetUserToken fetch a UserToken object by the specified token
func (h *RethinkDBDriver) GetUserToken(token string) (ut *models.UserToken, err error) {
	var res *r.Cursor
	res, err = r.Table(userTokenTableInit.TableName).
		GetAllByIndex("Token", token).
		Limit(1).
		CoerceTo("array").
		Run(h.conn)

	if err != nil {
		return ut, err
	}

	defer res.Close()

	rdata := map[string]interface{}{}

	if res.Next(&rdata) {
		ut = &models.UserToken{}
		err = convertFromRethinkDB(rdata, &ut)
		return ut, err
	}

	return ut, fmt.Errorf("not found")
}

// InvalidateUserTokens removes all user tokens that had been already expired
func (h *RethinkDBDriver) InvalidateUserTokens() (int, error) {
	wr, err := r.Table(userTokenTableInit.TableName).
		Filter(r.Row.Field("Expiration").Lt(time.Now())).
		Delete().
		RunWrite(h.conn)

	if err != nil {
		return 0, err
	}

	return wr.Deleted, nil
}
