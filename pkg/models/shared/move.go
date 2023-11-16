// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CurrentMark string

const (
	CurrentMarkX CurrentMark = "X"
	CurrentMarkO CurrentMark = "O"
)

func (e CurrentMark) ToPointer() *CurrentMark {
	return &e
}

func (e *CurrentMark) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = CurrentMark(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CurrentMark: %v", v)
	}
}

type Grid struct {
	Cells *string
}

func (o *Grid) GetCells() *string {
	if o == nil {
		return nil
	}
	return o.Cells
}

type MoveGrid struct {
	Cells *string
}

func (o *MoveGrid) GetCells() *string {
	if o == nil {
		return nil
	}
	return o.Cells
}

type MoveSchemasStartingMark string

const (
	MoveSchemasStartingMarkX MoveSchemasStartingMark = "X"
	MoveSchemasStartingMarkO MoveSchemasStartingMark = "O"
)

func (e MoveSchemasStartingMark) ToPointer() *MoveSchemasStartingMark {
	return &e
}

func (e *MoveSchemasStartingMark) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = MoveSchemasStartingMark(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MoveSchemasStartingMark: %v", v)
	}
}

type MoveAfterState struct {
	Grid         *MoveGrid
	StartingMark *MoveSchemasStartingMark
}

func (o *MoveAfterState) GetGrid() *MoveGrid {
	if o == nil {
		return nil
	}
	return o.Grid
}

func (o *MoveAfterState) GetStartingMark() *MoveSchemasStartingMark {
	if o == nil {
		return nil
	}
	return o.StartingMark
}

type MoveSchemasGrid struct {
	Cells *string
}

func (o *MoveSchemasGrid) GetCells() *string {
	if o == nil {
		return nil
	}
	return o.Cells
}

type MoveStartingMark string

const (
	MoveStartingMarkX MoveStartingMark = "X"
	MoveStartingMarkO MoveStartingMark = "O"
)

func (e MoveStartingMark) ToPointer() *MoveStartingMark {
	return &e
}

func (e *MoveStartingMark) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = MoveStartingMark(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MoveStartingMark: %v", v)
	}
}

type BeforeState struct {
	Grid         *MoveSchemasGrid
	StartingMark *MoveStartingMark
}

func (o *BeforeState) GetGrid() *MoveSchemasGrid {
	if o == nil {
		return nil
	}
	return o.Grid
}

func (o *BeforeState) GetStartingMark() *MoveStartingMark {
	if o == nil {
		return nil
	}
	return o.StartingMark
}

type Mark string

const (
	MarkX Mark = "X"
	MarkO Mark = "O"
)

func (e Mark) ToPointer() *Mark {
	return &e
}

func (e *Mark) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = Mark(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mark: %v", v)
	}
}

type PossibleMoves struct {
	AfterState  *MoveAfterState
	BeforeState *BeforeState
	CellIndex   *int64
	Mark        *Mark
}

func (o *PossibleMoves) GetAfterState() *MoveAfterState {
	if o == nil {
		return nil
	}
	return o.AfterState
}

func (o *PossibleMoves) GetBeforeState() *BeforeState {
	if o == nil {
		return nil
	}
	return o.BeforeState
}

func (o *PossibleMoves) GetCellIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.CellIndex
}

func (o *PossibleMoves) GetMark() *Mark {
	if o == nil {
		return nil
	}
	return o.Mark
}

type StartingMark string

const (
	StartingMarkX StartingMark = "X"
	StartingMarkO StartingMark = "O"
)

func (e StartingMark) ToPointer() *StartingMark {
	return &e
}

func (e *StartingMark) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = StartingMark(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StartingMark: %v", v)
	}
}

type Winner string

const (
	WinnerX Winner = "X"
	WinnerO Winner = "O"
)

func (e Winner) ToPointer() *Winner {
	return &e
}

func (e *Winner) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "X":
		fallthrough
	case "O":
		*e = Winner(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Winner: %v", v)
	}
}

type AfterState struct {
	CurrentMark    *CurrentMark
	GameNotStarted *bool
	GameOver       *bool
	Grid           *Grid
	PossibleMoves  []PossibleMoves
	StartingMark   *StartingMark
	Tie            *bool
	Winner         *Winner
	WinningCells   []int64
}

func (o *AfterState) GetCurrentMark() *CurrentMark {
	if o == nil {
		return nil
	}
	return o.CurrentMark
}

func (o *AfterState) GetGameNotStarted() *bool {
	if o == nil {
		return nil
	}
	return o.GameNotStarted
}

func (o *AfterState) GetGameOver() *bool {
	if o == nil {
		return nil
	}
	return o.GameOver
}

func (o *AfterState) GetGrid() *Grid {
	if o == nil {
		return nil
	}
	return o.Grid
}

func (o *AfterState) GetPossibleMoves() []PossibleMoves {
	if o == nil {
		return nil
	}
	return o.PossibleMoves
}

func (o *AfterState) GetStartingMark() *StartingMark {
	if o == nil {
		return nil
	}
	return o.StartingMark
}

func (o *AfterState) GetTie() *bool {
	if o == nil {
		return nil
	}
	return o.Tie
}

func (o *AfterState) GetWinner() *Winner {
	if o == nil {
		return nil
	}
	return o.Winner
}

func (o *AfterState) GetWinningCells() []int64 {
	if o == nil {
		return nil
	}
	return o.WinningCells
}

type Move struct {
	AfterState []AfterState
}

func (o *Move) GetAfterState() []AfterState {
	if o == nil {
		return nil
	}
	return o.AfterState
}
