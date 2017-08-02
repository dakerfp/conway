// Copyright (c) 2017, Daker Pinheiro
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//     * Redistributions of source code must retain the above copyright
//       notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright
//       notice, this list of conditions and the following disclaimer in the
//       documentation and/or other materials provided with the distribution.
//     * The names of its contributors may be used to endorse or promote products
//       derived from this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"flag"
	"fmt"
	"image"
	"time"
)

type Point image.Point

func Advance(board map[Point]bool) map[Point]bool {
	neighbours := make(map[Point]int)
	for p, _ := range board {
		// cardinal
		neighbours[Point{p.X, p.Y + 1}] += 1
		neighbours[Point{p.X, p.Y - 1}] += 1
		neighbours[Point{p.X + 1, p.Y}] += 1
		neighbours[Point{p.X - 1, p.Y}] += 1
		// diagonal
		neighbours[Point{p.X + 1, p.Y + 1}] += 1
		neighbours[Point{p.X + 1, p.Y - 1}] += 1
		neighbours[Point{p.X - 1, p.Y + 1}] += 1
		neighbours[Point{p.X - 1, p.Y - 1}] += 1
	}
	for p, n := range neighbours {
		_, alive := board[p]
		switch {
		case !alive && n == 3:
			board[p] = true
		case !alive:
			delete(board, p)
		case n < 2 || n > 3:
			delete(board, p)
		default:
			board[p] = true
		}
	}
	return board
}

var (
	blink = map[Point]bool{
		Point{-1, 0}: true,
		Point{0, 0}:  true,
		Point{1, 0}:  true,
	}
	glider = map[Point]bool{
		Point{-1, 0}: true,
		Point{0, 0}:  true,
		Point{1, 0}:  true,
		Point{1, 1}:  true,
		Point{0, 2}:  true,
	}
)

func writeGameWindow(board map[Point]bool, rect image.Rectangle) {
	for y := rect.Min.Y; y <= rect.Max.Y; y++ {
		for x := rect.Min.X; x <= rect.Max.X; x++ {
			if board[Point{x, y}] {
				fmt.Print("■")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	period := flag.Duration("period", time.Second/10, "time between animation steps")
	flag.Parse()
	board := glider
	timer := time.NewTicker(*period)
	for _ = range timer.C {
		fmt.Println("\033c") // Clear screen character
		writeGameWindow(board, image.Rect(-20, -20, 20, 20))
		board = Advance(board)
	}
}
