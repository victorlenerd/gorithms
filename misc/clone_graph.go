package misc

type UndirectedGraphNode struct {
	Label     int
	Neighbors []UndirectedGraphNode
}

func clone(startNode *UndirectedGraphNode) UndirectedGraphNode {

	cloneMap := map[int]UndirectedGraphNode{startNode.Label: *startNode}

	queue := make([]UndirectedGraphNode, 0)
	queue = append(queue, *startNode)

	for len(queue) > 0 {
		node := queue[0]

		for _, currentNeighbor := range node.Neighbors {
			neighbor, ok := cloneMap[currentNeighbor.Label]

			if !ok {
				cloneMap[currentNeighbor.Label] = node
				queue = append(queue, currentNeighbor)
			}

			neighbor.Neighbors = append(neighbor.Neighbors, cloneMap[currentNeighbor.Label])
		}

		queue = queue[1:]
	}

	return cloneMap[startNode.Label]
}
