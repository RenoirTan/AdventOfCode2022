package aoc2022

import (
	"errors"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Problem04 struct {
    Pairs [][2]Int64RangeInclusive
}

type Solution04 struct{}

type Day04 struct{}

func (day *Day04) BuildProblem(ctx *Context) (Problem, error) {
    raw := ctx.SplitLines()
    pairs := make([][2]Int64RangeInclusive, len(raw))
    for i, line := range raw {
        raw_pair := strings.Split(line, ",")
        pair_slice := lo.Map(raw_pair, func(r string, index int) Int64RangeInclusive {
            splat := strings.Split(r, "-")
            lower, err := strconv.Atoi(splat[0])
            if err != nil { panic(err) }
            upper, err := strconv.Atoi(splat[1])
            if err != nil { panic(err) }
            return Int64RangeInclusive{ int64(lower), int64(upper) }
        })
        pairs[i][0] = pair_slice[0]
        pairs[i][1] = pair_slice[1]
    }
    return &Problem04{pairs}, nil
}

func (day *Day04) BuildSolution(ctx *Context, problem Problem) (Solution, error) {
    return &Solution04{}, nil
}

func (sol *Solution04) P1(ctx *Context, problem Problem) (any, error) {
    p04 := TypeCast[Problem04](problem)
    if p04 == nil {
        return 0, errors.New("Bruh")
    }
    total := lo.SumBy(p04.Pairs, func(pair [2]Int64RangeInclusive) int64 {
        if pair[0].IsSubsetOf(&pair[1]) || pair[1].IsSubsetOf(&pair[0]) {
            return 1
        } else {
            return 0
        }
    })
    return total, nil
}

func (sol *Solution04) P2(ctx *Context, problem Problem) (any, error) {
    p04 := TypeCast[Problem04](problem)
    if p04 == nil {
        return 0, errors.New("Bruh")
    }
    total := lo.SumBy(p04.Pairs, func(pair [2]Int64RangeInclusive) int64 {
        if pair[0].Intersection(&pair[1]) != nil {
            return 1
        } else {
            return 0
        }
    })
    return total, nil
}