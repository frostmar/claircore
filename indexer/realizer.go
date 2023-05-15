package indexer

import (
	"context"

	"github.com/quay/claircore"
)

// Realizer is responsible for downloading a layer, uncompressing
// if necessary, and making the uncompressed tar contents available for
// reading.
type Realizer interface {
	Realize(context.Context, []*claircore.Layer) error
	Close() error
}

// FetchArena does coordination and global refcounting.
type FetchArena interface {
	Realizer(context.Context) Realizer
	Close(context.Context) error
}