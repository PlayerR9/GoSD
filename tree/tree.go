package tree

import (
	"errors"
	"iter"

	"github.com/PlayerR9/GoSD/pkg"
)

// Tree is the tree data structure.
type Tree[T interface {
	Child() iter.Seq[T]

	BackwardChild() iter.Seq[T]

	TreeNoder
}] struct {
	// root is the root node of the tree.
	root T
}

// String implements the fmt.Stringer interface.
func (t Tree[T]) String() string {
	trav := PrintFn[T]()

	info, err := ApplyDFS(&t, trav)
	if err != nil {
		pkg.Throw(err)
	}

	return info.String()
}

// DeepCopy implements the pkg.Type interface.
func (t *Tree[T]) DeepCopy() pkg.Type {
	if t == nil {
		return nil
	}

	root_copy := t.root.DeepCopy()

	tmp, ok := root_copy.(T)
	pkg.ThrowIf(!ok, pkg.NewInvalidState("root_copy", errors.New("invalid type")))

	return &Tree[T]{
		root: tmp,
	}
}

// Ensure implements the pkg.Type interface.
func (t *Tree[T]) Ensure() {
	pkg.ThrowIf(t == nil, pkg.NewInvalidState("t", pkg.NewNilValue()))

	t.root.Ensure()
}

// Clean implements the pkg.Type interface.
func (t *Tree[T]) Clean() {
	if t == nil {
		return
	}

	t.root.Clean()
}

// Equals implements the pkg.Type interface.
//
// Two trees are equal if they have the same root node.
func (t *Tree[T]) Equals(other pkg.Type) bool {
	pkg.Ensure(false, t)
	pkg.Ensure(false, other)

	switch other := other.(type) {
	case *Tree[T]:
		return t.root.Equals(other.root)
	default:
		return false
	}
}

// NewTree creates a new tree with the given root node.
//
// Parameters:
//   - root: The root node of the tree.
//
// Returns:
//   - *Tree[T]: The new tree. Never returns nil.
func NewTree[T interface {
	Child() iter.Seq[T]
	BackwardChild() iter.Seq[T]

	TreeNoder
}](root T) *Tree[T] {
	return &Tree[T]{
		root: root,
	}
}

// Root returns the root node of the tree.
//
// Returns:
//   - T: The root node of the tree.
func (t Tree[T]) Root() T {
	return t.root
}
