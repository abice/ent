// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dsl

// TraversalBuilder allows for an empty traversal to be .
type TraversalBuilder struct {
	*Traversal
}

// NewTraversalBuilder returns a traversal with no beginning token so that we can prepend on when
// it's known what starting point to use.
func NewTraversalBuilder() *TraversalBuilder {
	return &TraversalBuilder{
		Traversal: &Traversal{},
	}
}

// Clone creates a deep copy of an existing traversal.
func (tb *TraversalBuilder) Clone() *TraversalBuilder {
	if tb == nil {
		return nil
	}
	return &TraversalBuilder{Traversal: &Traversal{append(make([]Node, 0, len(tb.nodes)), tb.nodes...)}}
}

// AsTraversal allows us to pass in a builder to predicate functions.
func (tb *TraversalBuilder) AsTraversal() *Traversal {
	if tb == nil {
		return nil
	}
	return tb.Traversal
}

// BuildG will build the traversal with a prepended `g` node attached.
func (tb *TraversalBuilder) BuildG() *Traversal {
	if len(tb.nodes) < 1 {
		tb.Add(G)
		return tb.Traversal
	}

	if tb.nodes[0] == G {
		return tb.Traversal
	}

	if tb.nodes[0] == Token("__") {
		tb.nodes[0] = G
		return tb.Traversal
	}

	tb.nodes = append([]Node{G}, tb.nodes...)

	return tb.Traversal
}

// BuildAnonymous will build the traversal starting from `__`.
func (tb *TraversalBuilder) BuildAnonymous() *Traversal {
	if len(tb.nodes) < 1 {
		tb.Add(Token("__"))
		return tb.Traversal
	}

	if tb.nodes[0] == Token("__") {
		return tb.Traversal
	}

	if tb.nodes[0] == G {
		tb.nodes[0] = Token("__")
		return tb.Traversal
	}

	tb.nodes = append([]Node{Token("__")}, tb.nodes...)

	return tb.Traversal
}
