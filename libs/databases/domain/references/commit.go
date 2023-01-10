package references

import (
	"time"

	"github.com/steve-care-software/rodkina/libs/cryptography/domain/hash"
	"github.com/steve-care-software/rodkina/libs/cryptography/domain/hashtrees"
)

type commit struct {
	hash      hash.Hash
	values    hashtrees.HashTree
	createdOn time.Time
	pParent   *hash.Hash
	mine      Mine
}

func createCommit(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
) Commit {
	return createCommitInternally(hash, values, createdOn, nil, nil)
}

func createCommitWithParent(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
	pParent *hash.Hash,
) Commit {
	return createCommitInternally(hash, values, createdOn, pParent, nil)
}

func createCommitWithMine(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
	mine Mine,
) Commit {
	return createCommitInternally(hash, values, createdOn, nil, mine)
}

func createCommitWithParentAndMine(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
	pParent *hash.Hash,
	mine Mine,
) Commit {
	return createCommitInternally(hash, values, createdOn, pParent, mine)
}

func createCommitInternally(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
	pParent *hash.Hash,
	mine Mine,
) Commit {
	out := commit{
		hash:      hash,
		values:    values,
		createdOn: createdOn,
		pParent:   pParent,
		mine:      mine,
	}

	return &out
}

// Hash returns the hash
func (obj *commit) Hash() hash.Hash {
	return obj.hash
}

// Values returns the values
func (obj *commit) Values() hashtrees.HashTree {
	return obj.values
}

// CreatedOn returns the creation time
func (obj *commit) CreatedOn() time.Time {
	return obj.createdOn
}

// HasParent returns true if there is a parent, false otherwise
func (obj *commit) HasParent() bool {
	return obj.pParent != nil
}

// Parent returns the parent, if any
func (obj *commit) Parent() *hash.Hash {
	return obj.pParent
}

// HasMine returns true if there is a mine, false otherwise
func (obj *commit) HasMine() bool {
	return obj.mine != nil
}

// Mine returns the mine, if any
func (obj *commit) Mine() Mine {
	return obj.mine
}
