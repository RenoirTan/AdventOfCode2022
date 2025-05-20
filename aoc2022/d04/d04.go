package aoc2022_d04

import (
	"errors"
	"strconv"
	"strings"

	"github.com/RenoirTan/AdventOfCode2022/aoc2022"
	"github.com/samber/lo"
)

type Problem04 struct {
    Pairs [][2]aoc2022.Int64RangeInclusive
}

type Solution04 struct{}

type Day04 struct{}

func (day *Day04) BuildProblem(ctx *aoc2022.Context) (aoc2022.Problem, error) {
    raw := ctx.SplitLines()
    pairs := make([][2]aoc2022.Int64RangeInclusive, len(raw))
    for i, line := range raw {
        raw_pair := strings.Split(line, ",")
        pair_slice := lo.Map(raw_pair, func(r string, index int) aoc2022.Int64RangeInclusive {
            splat := strings.Split(r, "-")
            lower, err := strconv.Atoi(splat[0])
            if err != nil { panic(err) }
            upper, err := strconv.Atoi(splat[1])
            if err != nil { panic(err) }
            return aoc2022.Int64RangeInclusive{ Lower: int64(lower), Upper: int64(upper) }
        })
        pairs[i][0] = pair_slice[0]
        pairs[i][1] = pair_slice[1]
    }
    return &Problem04{pairs}, nil
}

func (day *Day04) BuildSolution(
    ctx *aoc2022.Context,
    problem aoc2022.Problem,
) (aoc2022.Solution, error) {
    return &Solution04{}, nil
}

func (sol *Solution04) P1(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p04 := aoc2022.TypeCast[Problem04](problem)
    if p04 == nil {
        return 0, errors.New("Bruh")
    }
    total := lo.SumBy(p04.Pairs, func(pair [2]aoc2022.Int64RangeInclusive) int64 {
        if pair[0].IsSubsetOf(&pair[1]) || pair[1].IsSubsetOf(&pair[0]) {
            return 1
        } else {
            return 0
        }
    })
    return total, nil
}

func (sol *Solution04) P2(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p04 := aoc2022.TypeCast[Problem04](problem)
    if p04 == nil {
        return 0, errors.New("Bruh")
    }
    total := lo.SumBy(p04.Pairs, func(pair [2]aoc2022.Int64RangeInclusive) int64 {
        if pair[0].Intersection(&pair[1]) != nil {
            return 1
        } else {
            return 0
        }
    })
    return total, nil
}