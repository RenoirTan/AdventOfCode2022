package aoc2022

import (
	"errors"

	"github.com/samber/lo"
)

func getPriority(r rune) int64 {
    if 'a' <= r  && r <= 'z' {
        return int64(r - 'a') + 1
    } else if 'A' <= r && r <= 'Z' {
        return int64(r - 'A') + 27
    } else {
        return 0
    }
}

type Problem03 struct {
    Rucksacks [][]rune
}

type Solution03 struct {}

type Day03 struct{}

func (day *Day03) BuildProblem(ctx *Context) (Problem, error) {
    raw := ctx.SplitLines()
    rucksacks := lo.Map(raw, func(item string, index int) []rune { return []rune(item) })
    return &Problem03{rucksacks}, nil
}

func (day *Day03) BuildSolution(ctx *Context, problem Problem) (Solution, error) {
    return &Solution03{}, nil
}

func (sol *Solution03) P1(ctx *Context, problem Problem) (int64, error) {
    p03 := TypeCast[Problem03](problem)
    if p03 == nil {
        return 0, errors.New("bruh")
    }
    total := int64(0)
    for _, rucksack := range p03.Rucksacks {
        halfway := len(rucksack) >> 1
        left := rucksack[:halfway]
        right := rucksack[halfway:]
        common := lo.Uniq(lo.Intersect(left, right))
        priorities := lo.SumBy(common, getPriority)
        total += priorities
    }
    return total, nil
}

func (sol *Solution03) P2(ctx *Context, problem Problem) (int64, error) {
    p03 := TypeCast[Problem03](problem)
    if p03 == nil {
        return 0, errors.New("bruh")
    }
    total := int64(0)
    for i := 0; i < len(p03.Rucksacks); i += 3 {
        r := p03.Rucksacks[i:i + 3]
        common := lo.Uniq(lo.Intersect(r[0], lo.Intersect(r[1], r[2])))
        if len(common) != 1 {
            return 0, errors.New("Expected common")
        }
        total += getPriority(common[0])
    }
    return total, nil
}