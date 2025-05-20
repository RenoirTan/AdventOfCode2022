package main

import (
	"fmt"

	"github.com/alecthomas/kong"

	"github.com/RenoirTan/AdventOfCode2022/aoc2022"
	aoc2022_d01 "github.com/RenoirTan/AdventOfCode2022/aoc2022/d01"
	aoc2022_d02 "github.com/RenoirTan/AdventOfCode2022/aoc2022/d02"
	aoc2022_d03 "github.com/RenoirTan/AdventOfCode2022/aoc2022/d03"
	aoc2022_d04 "github.com/RenoirTan/AdventOfCode2022/aoc2022/d04"
	aoc2022_d05 "github.com/RenoirTan/AdventOfCode2022/aoc2022/d05"
)

func GetDay(d uint64) aoc2022.Day {
    return []aoc2022.Day{
        nil,
        &aoc2022_d01.Day01{},
        &aoc2022_d02.Day02{},
        &aoc2022_d03.Day03{},
        &aoc2022_d04.Day04{},
        &aoc2022_d05.Day05{},
    }[d]
}

var CLI struct {
    Day int `arg:"" help:"Day to solve"`
    FilePath string `arg:"" help:"Path to input file"`
    Part *int `short:"p" help:"Which part of the day to solve"`
}

func main() {
    kong.Parse(&CLI)
    context := aoc2022.ContextDefault()
    context.OnDay(CLI.Day)
    context.WithInputFromPath(CLI.FilePath)
    day := GetDay(uint64(context.Day))
    problem, err := day.BuildProblem(&context)
    if err != nil {
        panic(err)
    }
    solution, err := day.BuildSolution(&context, problem)
    if err != nil {
        panic(err)
    }
    if CLI.Part == nil || *CLI.Part == 1 {
        answer, err := solution.P1(&context, problem)
        if err != nil {
            panic(err)
        } else {
            fmt.Print("Part 1: ")
            fmt.Println(answer)
        }
    }
    if CLI.Part == nil || *CLI.Part == 2 {
        answer, err := solution.P2(&context, problem)
        if err != nil {
            panic(err)
        } else {
            fmt.Print("Part 2: ")
            fmt.Println(answer)
        }
    }
}