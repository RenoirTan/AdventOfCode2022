package aoc2022_d09

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/RenoirTan/AdventOfCode2022/aoc2022"
	"github.com/golang/geo/r2"
	"github.com/samber/lo"
)

func ArePointsTouching(head r2.Point, tail r2.Point) bool {
    return math.Abs(head.X - tail.X) <= 1 && math.Abs(head.Y - tail.Y) <= 1
}

func PullTailTowardsHead(head r2.Point, tail r2.Point) r2.Point {
    if ArePointsTouching(head, tail) {
        return tail
    }

    diagonalDiff := head.Sub(tail)
    absDiagonalDiff := r2.Point{X: math.Abs(diagonalDiff.X), Y: math.Abs(diagonalDiff.Y)}
    minDiagonalDiff := 0.0
    if absDiagonalDiff.X == absDiagonalDiff.Y && absDiagonalDiff.X != 0 {
        minDiagonalDiff = absDiagonalDiff.X - 1
    } else {
        minDiagonalDiff = min(absDiagonalDiff.X, absDiagonalDiff.Y)
    }
    diagonalDiff = r2.Point{
        X: minDiagonalDiff * aoc2022.NormalizeScalar(diagonalDiff.X),
        Y: minDiagonalDiff * aoc2022.NormalizeScalar(diagonalDiff.Y),
    }
    tail = tail.Add(diagonalDiff)

    if ArePointsTouching(head, tail) {
        return tail
    }

    straightDiff := head.Sub(tail)
    straightDiff = r2.Point{
        X: max(math.Abs(straightDiff.X) - 1, 0) * aoc2022.NormalizeScalar(straightDiff.X),
        Y: max(math.Abs(straightDiff.Y) - 1, 0) * aoc2022.NormalizeScalar(straightDiff.Y),
    }

    return tail.Add(straightDiff)
}

type Instruction struct {
    Direction rune
    Distance int64
}

func (instruction *Instruction) AsNormalizedVector() r2.Point {
    switch (instruction.Direction) {
    case 'U':
        return r2.Point{X: 0, Y: -1}
    case 'D':
        return r2.Point{X: 0, Y: 1}
    case 'L':
        return r2.Point{X: -1, Y: 0}
    case 'R':
        return r2.Point{X: 1, Y: 0}
    default:
        return r2.Point{X: 0, Y: 0}
    }
}

func (instruction *Instruction) AsVector() r2.Point {
    return instruction.AsNormalizedVector().Mul(float64(instruction.Distance))
}

type Problem09 struct {
    Instructions []Instruction
}

type Chain struct {
    Knots []r2.Point
}

func MakeChain(length int) Chain {
    knots := make([]r2.Point, length)
    for i := range length {
        knots[i] = r2.Point{X: 0, Y: 0}
    }
    return Chain{knots}
}

func (chain *Chain) PullOnce(d r2.Point, visited *map[r2.Point]bool) {
    n := len(chain.Knots)
    if n < 2 {
        return
    }

    chain.Knots[0] = chain.Knots[0].Add(d)

    for i := 1; i < n; i++ {
        chain.Knots[i] = PullTailTowardsHead(chain.Knots[i - 1], chain.Knots[i])
    }
    (*visited)[chain.Knots[n - 1]] = true
}

func (chain Chain) String() string {
    list := strings.Join(lo.Map(chain.Knots, func(point r2.Point, index int) string {
        return fmt.Sprintf("(%d, %d)", int64(point.X), int64(point.Y))
    }), ", ")
    return "[" + list + "]"
}

type Solution09 struct {}

type Day09 struct {}

func (day *Day09) BuildProblem(ctx *aoc2022.Context) (aoc2022.Problem, error) {
    lines := ctx.SplitLines()
    instructions := make([]Instruction, len(lines))
    for i, line := range lines {
        direction := rune(line[0])
        distance, err := strconv.Atoi(line[2:])
        if err != nil {
            return nil, err
        }
        instructions[i] = Instruction{direction, int64(distance)}
    }
    return &Problem09{instructions}, nil
}

func (day *Day09) BuildSolution(
    ctx *aoc2022.Context,
    problem aoc2022.Problem,
) (aoc2022.Solution, error) {
    sol := &Solution09{}
    return sol, nil
}

func (sol *Solution09) Solve(ctx *aoc2022.Context, p09 *Problem09, length int) (any, error) {
    visited := map[r2.Point]bool{{X: 0, Y: 0}: true}
    chain := MakeChain(length)
    for _, instruction := range p09.Instructions {
        direction := instruction.AsNormalizedVector()
        for i := range instruction.Distance {
            _ = i // hence, the stupidity of go
            chain.PullOnce(direction, &visited)
        }
    }
    return len(visited), nil
}

func (sol *Solution09) P1(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p09 := aoc2022.TypeCast[Problem09](problem)
    if p09 == nil {
        return nil, errors.New("Bruh")
    }
    return sol.Solve(ctx, p09, 2)
}

func (sol *Solution09) P2(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p09 := aoc2022.TypeCast[Problem09](problem)
    if p09 == nil {
        return nil, errors.New("Bruh")
    }
    return sol.Solve(ctx, p09, 10)
}