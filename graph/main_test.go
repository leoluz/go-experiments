package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestAddEdge(t *testing.T) {
	t.Run("", func(t *testing.T) {
		// given
		t.Parallel()
		graph := &ItemGraph{}
		n1 := &Node{"leo"}
		n2 := &Node{"luz"}

		// when
		//graph.AddNode(n1)
		//graph.AddNode(n2)
		graph.AddEdge(n1, n2)

		// then
		assert.Equal(t, "bla", graph.String())
	})
}
