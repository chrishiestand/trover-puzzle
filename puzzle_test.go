package puzzle

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolveTrivial(t *testing.T) {
	DoneState := gameState{true, true, true,
		true, true, true,
		true, true, true}
	solution, err := solve(DoneState)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(solution)
}

func TestSolve1(t *testing.T) {
	aState := gameState{false, false, false,
		false, false, false,
		false, false, false}
	solution, err := solve(aState)
	if err != nil {
		t.Error(err)
	}
	if len(solution) < 1 {
		t.Error("solution cannot be empty")
	}
	fmt.Println(solution)
}

func TestUnsolvable(t *testing.T) {
	// This game state was taken from the game
	aState := gameState{false, false, true,
		true, true, true,
		true, true, true}
	_, err := solve(aState)
	if err == nil {
		t.Error("Expected no solution!")
	}
}

func TestMoves0(t *testing.T) {
	doneState := gameState{true, true, true,
		true, true, true,
		true, true, true}
	move, _ := genMove(0)
	expected := gameState{false, false, true, false, false, true, true, false, true}
	afterState := applyMove(doneState, move)
	if !reflect.DeepEqual(expected, afterState) {
		t.Errorf("State after move is not what was expected:\n%t\n%t", expected, afterState)

	}
}

func TestMoves1(t *testing.T) {
	doneState := gameState{true, false, true,
		false, true, false,
		true, false, true}
	move, _ := genMove(1)
	expected := gameState{false, true, false, false, false, false, false, false, false}
	afterState := applyMove(doneState, move)
	if !reflect.DeepEqual(expected, afterState) {
		t.Errorf("State after move is not what was expected:\n%t\n%t", expected, afterState)

	}
}
