package torrent

import (
	"testing"
	"unsafe"
)

func TestPieceSize(t *testing.T) {
	t.Log("[]*File", unsafe.Sizeof([]*File(nil)))
	t.Log("Piece", unsafe.Sizeof(Piece{}))
	t.Log("map[*peer]struct{}", unsafe.Sizeof(map[*peer]struct{}(nil)))
}