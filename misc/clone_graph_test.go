package misc

import (
	"testing"
)

func TestCloneGraph(t *testing.T) {

	t.Log("Clone graph")
	{
		t.Logf("\t Should create a clone of graph")
		{
			node1 := UndirectedGraphNode{Label: 1, Neighbors: []UndirectedGraphNode{}}
			node2 := UndirectedGraphNode{Label: 2, Neighbors: []UndirectedGraphNode{}}
			node3 := UndirectedGraphNode{Label: 3, Neighbors: []UndirectedGraphNode{}}
			node4 := UndirectedGraphNode{Label: 4, Neighbors: []UndirectedGraphNode{}}

			node2.Neighbors = append(node2.Neighbors, node1, node3)
			node3.Neighbors = append(node3.Neighbors, node2, node4)
			node4.Neighbors = append(node4.Neighbors, node1, node3)
			node1.Neighbors = append(node1.Neighbors, node2, node4)

			clonedGraph := clone(&node1)

			if len(clonedGraph.Neighbors) < 1 {
				t.Fatalf("\t Node one does not have neighbors")
			}

			if clonedGraph.Neighbors[0].Label != node2.Label {
				t.Fatalf("\t\t Node  %v does not match  %v",
					clonedGraph.Neighbors[0].Label, node2.Label)
			}

			if clonedGraph.Neighbors[0].Neighbors[1].Label != node3.Label {
				t.Fatalf("\t\t Node  %v does not match  %v",
					clonedGraph.Neighbors[0].Label, node2.Label)
			}

			if clonedGraph.Neighbors[1].Label != node4.Label {
				t.Fatalf("\t\t Node  %v does not match  %v",
					clonedGraph.Neighbors[0].Label, node2.Label)
			}
		}
	}
}
