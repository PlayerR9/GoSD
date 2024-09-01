package tree

import (
	"github.com/PlayerR9/GoSD/pkg"
)

// TreeNoder is the interface that describes the behavior of a tree node.
type TreeNoder interface {
	// IsLeaf is a method that checks whether the node is a leaf node.
	//
	// Returns:
	//   - bool: True if the node is a leaf node, false otherwise.
	IsLeaf() bool

	comparable

	pkg.Type
}
