package puzzle

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type gameState []bool

var unixTime = time.Now().UnixNano()
var seed = rand.NewSource(unixTime)
var prg = rand.New(seed)

func solve(initState gameState) ([]int, error) {
	path := []int{}
	var state = make(gameState, 9)

	copy(state, initState)

	for i := 0; i < 9999999; i++ {
		// fmt.Println("state report:", state)

		if isDoneState(state) {
			fmt.Println("done")

			return path, nil
		}
		nextMove := genNextMove()
		path = append(path, nextMove)
		// fmt.Println("nextMove", nextMove)
		toggled, err := genMove(nextMove)
		if err != nil {
			err2 := fmt.Errorf("no rule found for nextMove %q", nextMove)
			return nil, err2
		}
		state = applyMove(state, toggled)
	}
	return nil, errors.New("Could not find a solution")
}

func applyMove(state gameState, toggled []int) gameState {
	for _, stateKey := range toggled {
		state[stateKey] = !state[stateKey]
	}
	return state
}

func genMove(move int) ([]int, error) {
	switch move {
	case 0:
		return []int{0, 1, 3, 4, 7}, nil
	case 1:
		return []int{0, 1, 2, 4, 6, 8}, nil
	case 2:
		return []int{0, 2, 3, 6}, nil
	case 3:
		return []int{0, 1, 2, 3}, nil
	case 4:
		return []int{3, 4, 5}, nil
	case 5:
		return []int{2, 5, 8}, nil
	case 6:
		return []int{6, 4, 2}, nil
	case 7:
		return []int{7, 4, 1}, nil
	case 8:
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, nil
	}

	err := fmt.Errorf("move %q not supported", move)
	return nil, err
}

func genNextMove() int {
	return prg.Intn(9)
}

func isDoneState(state gameState) bool {

	for _, piece := range state {

		if !piece {
			return false
		}
	}
	return true
}
