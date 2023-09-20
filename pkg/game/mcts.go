package game

import (
	"math"
	"math/rand"
)

// MctsMove returns the best move for the given game state using Monte Carlo Tree Search.
func MctsMove(game Othello, iterations int) [2]int {
	root := new_node(nil, [2]int{}, game.State, game.GetValidMoves())

	for i := 0; i < iterations; i++ {
		node := root
		simulation := game.DeepCopy()
		// SELECT
		for len(node.unexplored) == 0 && len(node.children) > 0 {
			node = node.select_child()
			simulation.MakeMove(node.position)
		}
		// EXPAND
		if len(node.unexplored) > 0 {
			rand_idx := rand.Intn(len(node.unexplored))
			explored_move := node.unexplored[rand_idx]
			explored_turn := simulation.State
			simulation.MakeMove(explored_move)
			// remove explored move from unexplored and add child to the tree
			node.unexplored = append(node.unexplored[:rand_idx], node.unexplored[rand_idx+1:]...)
			child := new_node(node, explored_move, explored_turn, simulation.GetValidMoves())
			node.children = append(node.children, child)
			node = child
		}
		// SIMULATE
		for simulation.State == BLACK_TURN || simulation.State == WHITE_TURN {
			possible_moves := simulation.GetValidMoves()
			simulation.MakeMove(possible_moves[rand.Intn(len(possible_moves))])
		}
		// BACKPROPAGATE
		for node != nil {
			node.visits++
			if simulation.State == DRAW {
				// do nothing because draw is neutral
			} else if (simulation.State == BLACK_WON) == (node.turn == BLACK_TURN) {
				node.wins++
			} else {
				node.wins--
			}
			node = node.parent
		}
	}
	return root.get_most_visited_position()
}

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
	best_uct := math.Inf(-1)
	ln_total := math.Log(2 * float64(n.visits))
	for child_idx, child := range n.children {
		uct := float64(child.wins)/float64(child.visits) + math.Sqrt(ln_total/float64(child.visits))
		if uct > best_uct {
			best_child_idx = child_idx
			best_uct = uct
		}
	}
	return n.children[best_child_idx]
}

func (n *Node) get_most_visited_position() [2]int {
	most_visited_child_idx := 0
	most_visits := 0
	for i, child := range n.children {
		if child.visits > most_visits {
			most_visited_child_idx = i
			most_visits = child.visits
		}
	}
	return n.children[most_visited_child_idx].position
}
