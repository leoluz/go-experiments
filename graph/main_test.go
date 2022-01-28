package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestAddEdge(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
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
	t.Run("test 2", func(t *testing.T) {
		// given
		t.Parallel()
		graph := &ItemGraph{}
		n1 := &Node{"hi"}
		n2 := &Node{"debug"}

		// when
		graph.AddNode(n1)
		graph.AddNode(n2)
		graph.AddEdge(n1, n2)

		// then
		assert.Equal(t, "hi debug", graph.String())
	})
}

func TestDebug(t *testing.T) {
	t.Run("debugging", func(t *testing.T) {
		// given
		t.Parallel()
		graph := &ItemGraph{}
		n1 := &Node{"debug"}
		n2 := &Node{"this"}

		// when
		graph.AddNode(n1)
		graph.AddNode(n2)
		graph.AddEdge(n1, n2)

		// then
		assert.Equal(t, "hi debug", graph.String())
	})
}
