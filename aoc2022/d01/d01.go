package aoc2022_d01

import (
	"errors"
	"strconv"

	"github.com/RenoirTan/AdventOfCode2022/aoc2022"
	"github.com/samber/lo"
)


type Problem01 struct {
    Elves [][]int64
}

type Solution01 struct{
    RationTotals []int64
}

type Day01 struct{}

func (day *Day01) BuildProblem(ctx *aoc2022.Context) (aoc2022.Problem, error) {
    lines := ctx.SplitLines()
    n_elves := 1
    for _, line := range lines {
        if len(line) == 0 {
            n_elves++
        }
    }
    elves := make([][]int64, n_elves)
    i := 0
    rations := make([]int64, 0)
    for _, line := range lines {
        if len(line) == 0 {
            elves[i] = rations
            i++
            rations = make([]int64, 0)
        } else {
            ration, err := strconv.Atoi(line)
            if err != nil {
                return nil, err
            }
            rations = append(rations, int64(ration))
        }
    }
    elves[i] = rations
    return &Problem01{elves}, nil
}

func (day *Day01) BuildSolution(
    ctx *aoc2022.Context,
    problem aoc2022.Problem,
) (aoc2022.Solution, error) {
    p01 := aoc2022.TypeCast[Problem01](problem)
    if p01 == nil {
        return nil, errors.New("bruh")
    }
    solution := &Solution01{}
    solution.sumRations(p01)
    return solution, nil
}

func (sol *Solution01) sumRations(p01 *Problem01) {
    if sol.RationTotals != nil {
        return
    }
    totals := make([]int64, len(p01.Elves))
    for elf, rations := range p01.Elves {
        calories := lo.Sum(rations)
        totals[elf] = calories
    }
    aoc2022.Sort(totals)
    sol.RationTotals = totals
}

func (sol *Solution01) P1(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    return sol.RationTotals[len(sol.RationTotals) - 1], nil
}

func (sol *Solution01) P2(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    return lo.Sum(sol.RationTotals[len(sol.RationTotals) - 3:]), nil
}