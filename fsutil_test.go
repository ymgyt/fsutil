package fsutil

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dammySyncCloser struct{}

func (d *dammySyncCloser) Sync() error {
	return errors.New("sync error")
}
func (d *dammySyncCloser) Close() error {
	return errors.New("close error")
}

func TestClose(t *testing.T) {
	var err error
	var d dammySyncCloser
	Close(&d, &err)
	assert.NotNil(t, err)

	orgErr := errors.New("orginal error")
	err = orgErr
	Close(&d, &err)
	assert.Equal(t, err, orgErr)
}

func TestSyncClose(t *testing.T) {
	var err error
	var d dammySyncCloser
	SyncClose(&d, &err)
	assert.NotNil(t, err)

	orgErr := errors.New("orginal error")
	err = orgErr
	SyncClose(&d, &err)
	assert.Equal(t, err, orgErr)
}
