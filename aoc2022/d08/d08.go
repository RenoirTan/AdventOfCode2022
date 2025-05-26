package aoc2022_d08

import (
	"errors"

	"github.com/RenoirTan/AdventOfCode2022/aoc2022"
	"github.com/samber/lo"
)

type Problem08 struct {
    trees [][]int64
    height int
    width int
}

type Solution08 struct {
    visible [][]bool
}

type Day08 struct {}

func (day *Day08) BuildProblem(ctx *aoc2022.Context) (aoc2022.Problem, error) {
    lines := ctx.SplitLines()
    height := len(lines)
    width := len(lines[0])
    trees := make([][]int64, len(lines))
    for y, line := range lines {
        row := make([]int64, len(line))
        for x, c := range line {
            row[x] = int64(c - '0')
        }
        trees[y] = row
    }
    return &Problem08{trees, height, width}, nil
}

func (day *Day08) BuildSolution(
    ctx *aoc2022.Context,
    problem aoc2022.Problem,
) (aoc2022.Solution, error) {
    p08 := aoc2022.TypeCast[Problem08](problem)
    if p08 == nil {
        return nil, errors.New("Bruh")
    }
    visible := make([][]bool, p08.height)
    for y := range p08.height {
        visible[y] = make([]bool, p08.width)
    }
    return &Solution08{visible}, nil
}

func visitTree(visible [][]bool, trees [][]int64, x int, y int, tallest int64) int64 {
    tree := trees[y][x]
    if tree > tallest {
        visible[y][x] = true
        return tree
    } else {
        return tallest
    }
}

func scenicScoreOf(p08 *Problem08, x, y int) int64 {
    score := int64(1)
    tree := p08.trees[y][x]
    nx, ny := 0, 0
    for ny = y - 1; ny >= 0; ny-- {
        if p08.trees[ny][x] >= tree {
            break
        }
    }
    score *= int64(y - max(ny, 0))
    for ny = y + 1; ny < p08.height; ny++ {
        if p08.trees[ny][x] >= tree {
            break
        }
    }
    score *= int64(min(ny, p08.height-1) - y)
    for nx = x - 1; nx >= 0; nx-- {
        if p08.trees[y][nx] >= tree {
            break
        }
    }
    score *= int64(x - max(nx, 0))
    for nx = x + 1; nx < p08.width; nx++ {
        if p08.trees[y][nx] >= tree {
            break
        }
    }
    score *= int64(min(nx, p08.width-1) - x)
    return score
}

func (sol *Solution08) P1(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p08 := aoc2022.TypeCast[Problem08](problem)
    if p08 == nil {
        return nil, errors.New("Bruh")
    }
    for y := range p08.height {
        lo.Reduce(lo.RangeWithSteps(0, p08.width, 1), func(tallest int64, x, _ int) int64 {
            return visitTree(sol.visible, p08.trees, x, y, tallest)
        }, -1)
        lo.Reduce(lo.RangeWithSteps(p08.width - 1, -1, -1), func(tallest int64, x, _ int) int64 {
            return visitTree(sol.visible, p08.trees, x, y, tallest)
        }, -1)
    }
    for x := range p08.width {
        lo.Reduce(lo.RangeWithSteps(0, p08.height, 1), func(tallest int64, y, _ int) int64 {
            return visitTree(sol.visible, p08.trees, x, y, tallest)
        }, -1)
        lo.Reduce(lo.RangeWithSteps(p08.height - 1, -1, -1), func(tallest int64, y, _ int) int64 {
            return visitTree(sol.visible, p08.trees, x, y, tallest)
        }, -1)
    }
    total := lo.SumBy(lo.Flatten(sol.visible), aoc2022.BToi)
    return total, nil
}

func (sol *Solution08) P2(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p08 := aoc2022.TypeCast[Problem08](problem)
    if p08 == nil {
        return nil, errors.New("Bruh")
    }
    scores := lo.CrossJoinBy2(
        lo.RangeFrom(1, p08.width-2),
        lo.RangeFrom(1, p08.height-2),
        func(x, y int) int64 {
            return scenicScoreOf(p08, x, y)
        },
    )
    return lo.Max(scores), nil
}