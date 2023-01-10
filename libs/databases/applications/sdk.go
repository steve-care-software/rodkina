package applications

import (
	"github.com/steve-care-software/rodkina/libs/cryptography/domain/hash"
	"github.com/steve-care-software/rodkina/libs/databases/domain/references"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	Now() (Application, error)
}

// Application represents the read application
type Application interface {
	Database
	Reference
	Content
}

// Reference represents the reference application
type Reference interface {
	ContentKeysByKind(context uint, kind uint) (references.ContentKeys, error)
	CommitByHash(context uint, hash hash.Hash) (references.Commit, error)
	Commits(context uint) (references.Commits, error)
}

// Database represents the database application
type Database interface {
	Exists(name string) (bool, error)
	New(name string) error
	Delete(name string) error
	Open(name string) (*uint, error)
	Cancel(context uint) error
	Commit(context uint) error
	Close(context uint) error
}

// Content represents the content application
type Content interface {
	Read(context uint, pointer references.Pointer) ([]byte, error)
	ReadByHash(context uint, hash hash.Hash) ([]byte, error)
	ReadAll(context uint, pointers []references.Pointer) ([][]byte, error)
	ReadAllByHashes(context uint, hashes []hash.Hash) ([][]byte, error)
	Write(context uint, hash hash.Hash, data []byte, kind uint) error
}
