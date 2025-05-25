package aoc2022_d06

import (
	"errors"

	"github.com/RenoirTan/AdventOfCode2022/aoc2022"
	"github.com/samber/lo"
)

type Problem06 struct {
    signal []rune
}

type Solution06 struct {}

type Day06 struct {}

func (day *Day06) BuildProblem(ctx *aoc2022.Context) (aoc2022.Problem, error) {
    signal := []rune(ctx.Input)
    return &Problem06{signal}, nil
}

func (day *Day06) BuildSolution(
    ctx *aoc2022.Context,
    problem aoc2022.Problem,
) (aoc2022.Solution, error) {
    return &Solution06{}, nil
}

func allUnique(window []rune, size int) bool {
    return len(window) == size && len(lo.Uniq(window)) == size
}

func solve(problem aoc2022.Problem, size int) (any, error) {
    p06 := aoc2022.TypeCast[Problem06](problem)
    if p06 == nil {
        return 0, errors.New("Bruh")
    }
    for i := range len(p06.signal) - size {
        if allUnique(p06.signal[i:i + size], size) {
            return i + size, nil
        }
    }
    return -1, nil
}

func (sol *Solution06) P1(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    return solve(problem, 4)
}

func (sol *Solution06) P2(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    return solve(problem, 14)
}