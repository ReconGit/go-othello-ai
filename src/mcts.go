package src

import (
	"github.com/huandu/go-clone/generic"
	"math"
	"math/rand"
)

type Node struct {
	parent     *Node
	position   [2]int
	turn       State
	unexplored [][2]int
	children   []*Node
	visits     int
	wins       int
}

func new_node(parent *Node, position [2]int, turn State, unexplored [][2]int) *Node {
	return &Node{
		parent:     parent,
		position:   position,
		turn:       turn,
		unexplored: unexplored,
		children:   []*Node{},
		visits:     0,
		wins:       0,
	}
}

func (n *Node) select_child() *Node {
	best_child_idx := 0
	best_score := math.Inf(-1)
	log_total := math.Log(2 * float64(n.visits))
	for child_idx, child := range n.children {
		score := float64(child.wins)/float64(child.visits) + math.Sqrt(log_total/float64(child.visits))
		if score > best_score {
			best_child_idx = child_idx
			best_score = score
		}
	}
	return n.children[best_child_idx]
}

func (n *Node) get_most_visited_position() [2]int {
	var most_visited *Node
	most_visits := 0
	for _, child := range n.children {
		if child.visits > most_visits {
			most_visited = child
			most_visits = child.visits
		}
	}
	return most_visited.position
}

func MctsMove(game Othello, iterations int) [2]int {
	root := new_node(nil, [2]int{}, game.state, game.GetValidMoves())
	// for interations
	for i := 0; i < iterations; i++ {
		node := root
		simulation := clone.Clone(game)
		// SELECT
		for len(node.unexplored) == 0 && len(node.children) > 0 {
			node = node.select_child()
			simulation.MakeMove(node.position)
		}
		// EXPAND
		if len(node.unexplored) > 0 {
			rand_idx := rand.Intn(len(node.unexplored))
			explored_move := node.unexplored[rand.Intn(len(node.unexplored))]
			explored_turn := simulation.state
			simulation.MakeMove(explored_move)
			// remove explored move from unexplored and add child to the tree
			node.unexplored = append(node.unexplored[:rand_idx], node.unexplored[rand_idx+1:]...)
			child := new_node(node, explored_move, explored_turn, simulation.GetValidMoves())
			node.children = append(node.children, child)
			node = child
		}
		// SIMULATE
		for simulation.state == BLACK_TURN || simulation.state == WHITE_TURN {
			possible_moves := simulation.GetValidMoves()
			simulation.MakeMove(possible_moves[rand.Intn(len(possible_moves))])
		}
		// BACKPROPAGATE
		for node != nil {
			node.visits++
			if simulation.state == DRAW {
				// do nothing because draw is neutral
			} else if (simulation.state == BLACK_WON) == (node.turn == BLACK_TURN) {
				node.wins++
			} else {
				node.wins--
			}
			node = node.parent
		}
	}
	return root.get_most_visited_position()
}
