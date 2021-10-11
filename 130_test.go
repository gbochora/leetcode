package main

import (
	"reflect"
	"testing"
)

func TestSurroundedRegions1(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'}}
	Solve(board)
	want := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'}}
	if !reflect.DeepEqual(want, board) {
		t.Errorf("got %v, want %v", board, want)
	}
}

func TestSurroundedRegions2(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X'}}
	Solve(board)
	want := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X'}}
	if !reflect.DeepEqual(want, board) {
		t.Errorf("got %v, want %v", board, want)
	}
}

func TestSurroundedRegions3(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'X', 'X', 'X'}}
	Solve(board)
	want := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'}}
	if !reflect.DeepEqual(want, board) {
		t.Errorf("got %v, want %v", board, want)
	}
}
