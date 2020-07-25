package db

import "fmt"

type Blob struct {
	Hash  string
	Owner int64
	Size  uint64
	Path  string
}

func (b *Blob) String() string {
	return fmt.Sprintf("Blob<%s, %d, %d, %s>", b.Hash, b.Owner, b.Size, b.Path)
}
