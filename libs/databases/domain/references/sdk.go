package references

import (
	"math/big"
	"time"

	"github.com/steve-care-software/rodkina/libs/cryptography/domain/hash"
	"github.com/steve-care-software/rodkina/libs/cryptography/domain/hashtrees"
)

const pointerSize = 8 * 2
const commitMinSize = 8 + 8 + hashtrees.MinHashtreeSize
const contentKeySize = hash.Size + pointerSize + 8 + hash.Size
const minReferenceSize = contentKeySize + commitMinSize

// NewAdapter creates a new adapter instance
func NewAdapter(
	miningValue byte,
) Adapter {
	contentKeysAdapter := NewContentKeysAdapter()
	commitsAdapter := NewCommitsAdapter(miningValue)
	builder := NewBuilder()
	return createAdapter(
		contentKeysAdapter,
		commitsAdapter,
		builder,
	)
}

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewCommitsAdapter creates a new commits adapter
func NewCommitsAdapter(
	miningValue byte,
) CommitsAdapter {
	adapter := NewCommitAdapter(miningValue)
	builder := NewCommitsBuilder()
	return createCommitsAdapter(adapter, builder)
}

// NewCommitsBuilder creates a new commits builder
func NewCommitsBuilder() CommitsBuilder {
	return createCommitsBuilder()
}

// NewCommitAdapter creates a new commit adapter
func NewCommitAdapter(
	miningValue byte,
) CommitAdapter {
	hashAdapter := hash.NewAdapter()
	hashTreeAdapter := hashtrees.NewAdapter()
	builder := NewCommitBuilder(miningValue)
	return createCommitAdapter(hashAdapter, hashTreeAdapter, builder)
}

// NewCommitBuilder creates a new commit builder
func NewCommitBuilder(
	miningValue byte,
) CommitBuilder {
	hashAdapter := hash.NewAdapter()
	return createCommitBuilder(
		hashAdapter,
		miningValue,
	)
}

// NewContentKeysAdapter creates a new content keys adapter
func NewContentKeysAdapter() ContentKeysAdapter {
	adapter := NewContentKeyAdapter()
	builder := NewContentKeysBuilder()
	return createContentKeysAdapter(adapter, builder)
}

// NewContentKeysBuilder creates a new content keys builder
func NewContentKeysBuilder() ContentKeysBuilder {
	return createContentKeysBuilder()
}

// NewContentKeyAdapter creates a new content key adapter
func NewContentKeyAdapter() ContentKeyAdapter {
	hashAdapter := hash.NewAdapter()
	pointerAdapter := NewPointerAdapter()
	builder := NewContentKeyBuilder()
	return createContentKeyAdapter(hashAdapter, pointerAdapter, builder)
}

// NewContentKeyBuilder createsa new content key builder
func NewContentKeyBuilder() ContentKeyBuilder {
	return createContentKeyBuilder()
}

// NewPointerAdapter creates a new pointer adapter
func NewPointerAdapter() PointerAdapter {
	builder := NewPointerBuilder()
	return createPointerAdapter(builder)
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	return createPointerBuilder()
}

// Adapter represents a reference adapter
type Adapter interface {
	ToContent(ins Reference) ([]byte, error)
	ToReference(content []byte) (Reference, error)
}

// Factory represents a reference factory
type Factory interface {
	Create() (Reference, error)
}

// Builder represents a reference builder
type Builder interface {
	Create() Builder
	WithContentKeys(contentKeys ContentKeys) Builder
	WithCommits(commits Commits) Builder
	Now() (Reference, error)
}

// Reference represents the reference
type Reference interface {
	ContentKeys() ContentKeys
	Commits() Commits
}

// CommitsAdapter represents a commits adapter
type CommitsAdapter interface {
	ToContent(ins Commits) ([]byte, error)
	ToCommits(content []byte) (Commits, error)
}

// CommitsBuilder represents commits builder
type CommitsBuilder interface {
	Create() CommitsBuilder
	WithList(list []Commit) CommitsBuilder
	Now() (Commits, error)
}

// Commits represents commits
type Commits interface {
	List() []Commit
	Latest() Commit
	Fetch(hash hash.Hash) (Commit, error)
}

// CommitAdapter represents a commit adapter
type CommitAdapter interface {
	ToContent(ins Commit) ([]byte, error)
	ToCommit(content []byte) (Commit, error)
}

// CommitBuilder represents a commit builder
type CommitBuilder interface {
	Create() CommitBuilder
	WithValues(values hashtrees.HashTree) CommitBuilder
	WithParent(parent hash.Hash) CommitBuilder
	WithProof(proof *big.Int) CommitBuilder
	CreatedOn(createdOn time.Time) CommitBuilder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Values() hashtrees.HashTree
	CreatedOn() time.Time
	HasParent() bool
	Parent() *hash.Hash
	HasMine() bool
	Mine() Mine
}

// Mine represents a mine
type Mine interface {
	Result() hash.Hash
	Proof() *big.Int
	Score() uint
}

// ContentKeysAdapter represents the content keys adapter
type ContentKeysAdapter interface {
	ToContent(ins ContentKeys) ([]byte, error)
	ToContentKeys(content []byte) (ContentKeys, error)
}

// ContentKeysBuilder represents a content keys builder
type ContentKeysBuilder interface {
	Create() ContentKeysBuilder
	WithList(list []ContentKey) ContentKeysBuilder
	Now() (ContentKeys, error)
}

// ContentKeys represents content keys
type ContentKeys interface {
	Next() int64
	List() []ContentKey
	ListByKind(kind uint) []ContentKey
	Fetch(hash hash.Hash) (ContentKey, error)
}

// ContentKeyAdapter represents the content key adapter
type ContentKeyAdapter interface {
	ToContent(ins ContentKey) ([]byte, error)
	ToContentKey(content []byte) (ContentKey, error)
}

// ContentKeyBuilder represents a content key builder
type ContentKeyBuilder interface {
	Create() ContentKeyBuilder
	WithHash(hash hash.Hash) ContentKeyBuilder
	WithKind(kind uint) ContentKeyBuilder
	WithContent(content Pointer) ContentKeyBuilder
	WithCommit(commit hash.Hash) ContentKeyBuilder
	Now() (ContentKey, error)
}

// ContentKey represents a content key
type ContentKey interface {
	Hash() hash.Hash
	Content() Pointer
	Kind() uint
	Commit() hash.Hash
}

// PointerAdapter represents the pointer adapter
type PointerAdapter interface {
	ToContent(ins Pointer) ([]byte, error)
	ToPointer(content []byte) (Pointer, error)
}

// PointerBuilder represents a pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithLength(length uint) PointerBuilder
	From(from uint) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	From() uint
	Length() uint
}
