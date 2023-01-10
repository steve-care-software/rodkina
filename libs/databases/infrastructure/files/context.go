package files

import (
	"os"

	"github.com/juju/fslock"
	"github.com/steve-care-software/rodkina/libs/databases/domain/contents"
	"github.com/steve-care-software/rodkina/libs/databases/domain/references"
)

type context struct {
	identifier  uint
	name        string
	pLock       *fslock.Lock
	pConn       *os.File
	reference   references.Reference
	dataOffset  uint
	contentList []contents.Content
}
