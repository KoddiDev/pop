// +build !sqlite

package pop

import (
	"errors"
)

func newSQLite(deets *ConnectionDetails) (dialect, error) {
	return nil, errors.New("sqlite3 support was not compiled into the binary")
}
