// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type MoveAfterStateCurrentMark string

const (
	MoveAfterStateCurrentMarkX MoveAfterStateCurrentMark = "X"
	MoveAfterStateCurrentMarkO MoveAfterStateCurrentMark = "O"
)

func (e MoveAfterStateCurrentMark) ToPointer() *MoveAfterStateCurrentMark {
	return &e
}

func (e *MoveAfterStateCurrentMark) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = MoveAfterStateCurrentMark(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MoveAfterStateCurrentMark: %v", v)
	}
}

type MoveAfterStateGrid struct {
	Cells *string
}

type MoveAfterStatePossibleMovesAfterStateGrid struct {
	Cells *string
}

type MoveAfterStatePossibleMovesAfterStateStartingMark string

const (
	MoveAfterStatePossibleMovesAfterStateStartingMarkX MoveAfterStatePossibleMovesAfterStateStartingMark = "X"
	MoveAfterStatePossibleMovesAfterStateStartingMarkO MoveAfterStatePossibleMovesAfterStateStartingMark = "O"
)

func (e MoveAfterStatePossibleMovesAfterStateStartingMark) ToPointer() *MoveAfterStatePossibleMovesAfterStateStartingMark {
	return &e
}

func (e *MoveAfterStatePossibleMovesAfterStateStartingMark) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = MoveAfterStatePossibleMovesAfterStateStartingMark(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MoveAfterStatePossibleMovesAfterStateStartingMark: %v", v)
	}
}

type MoveAfterStatePossibleMovesAfterState struct {
	Grid         *MoveAfterStatePossibleMovesAfterStateGrid
	StartingMark *MoveAfterStatePossibleMovesAfterStateStartingMark
}

type MoveAfterStatePossibleMovesBeforeStateGrid struct {
	Cells *string
}

type MoveAfterStatePossibleMovesBeforeStateStartingMark string

const (
	MoveAfterStatePossibleMovesBeforeStateStartingMarkX MoveAfterStatePossibleMovesBeforeStateStartingMark = "X"
	MoveAfterStatePossibleMovesBeforeStateStartingMarkO MoveAfterStatePossibleMovesBeforeStateStartingMark = "O"
)

func (e MoveAfterStatePossibleMovesBeforeStateStartingMark) ToPointer() *MoveAfterStatePossibleMovesBeforeStateStartingMark {
	return &e
}

func (e *MoveAfterStatePossibleMovesBeforeStateStartingMark) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = MoveAfterStatePossibleMovesBeforeStateStartingMark(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MoveAfterStatePossibleMovesBeforeStateStartingMark: %v", v)
	}
}

type MoveAfterStatePossibleMovesBeforeState struct {
	Grid         *MoveAfterStatePossibleMovesBeforeStateGrid
	StartingMark *MoveAfterStatePossibleMovesBeforeStateStartingMark
}

type MoveAfterStatePossibleMovesMark string

const (
	MoveAfterStatePossibleMovesMarkX MoveAfterStatePossibleMovesMark = "X"
	MoveAfterStatePossibleMovesMarkO MoveAfterStatePossibleMovesMark = "O"
)

func (e MoveAfterStatePossibleMovesMark) ToPointer() *MoveAfterStatePossibleMovesMark {
	return &e
}

func (e *MoveAfterStatePossibleMovesMark) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = MoveAfterStatePossibleMovesMark(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MoveAfterStatePossibleMovesMark: %v", v)
	}
}

type MoveAfterStatePossibleMoves struct {
	AfterState  *MoveAfterStatePossibleMovesAfterState
	BeforeState *MoveAfterStatePossibleMovesBeforeState
	CellIndex   *int64
	Mark        *MoveAfterStatePossibleMovesMark
}

type MoveAfterStateStartingMark string

const (
	MoveAfterStateStartingMarkX MoveAfterStateStartingMark = "X"
	MoveAfterStateStartingMarkO MoveAfterStateStartingMark = "O"
)

func (e MoveAfterStateStartingMark) ToPointer() *MoveAfterStateStartingMark {
	return &e
}

func (e *MoveAfterStateStartingMark) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = MoveAfterStateStartingMark(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MoveAfterStateStartingMark: %v", v)
	}
}

type MoveAfterStateWinner string

const (
	MoveAfterStateWinnerX MoveAfterStateWinner = "X"
	MoveAfterStateWinnerO MoveAfterStateWinner = "O"
)

func (e MoveAfterStateWinner) ToPointer() *MoveAfterStateWinner {
	return &e
}

func (e *MoveAfterStateWinner) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = MoveAfterStateWinner(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MoveAfterStateWinner: %v", v)
	}
}

type MoveAfterState struct {
	CurrentMark    *MoveAfterStateCurrentMark
	GameNotStarted *bool
	GameOver       *bool
	Grid           *MoveAfterStateGrid
	PossibleMoves  []MoveAfterStatePossibleMoves
	StartingMark   *MoveAfterStateStartingMark
	Tie            *bool
	Winner         *MoveAfterStateWinner
	WinningCells   []int64
}

// Move - A Move containing the before and after GameStates.
type Move struct {
	AfterState []MoveAfterState
}
