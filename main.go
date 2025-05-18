package main

import (
	"fmt"

	"github.com/alecthomas/kong"

	"github.com/RenoirTan/AdventOfCode2022/aoc2022"
)

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
    day := aoc2022.GetDay(uint64(context.Day))
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
            fmt.Printf("Part 1: %d\n", answer)
        }
    }
    if CLI.Part == nil || *CLI.Part == 2 {
        answer, err := solution.P2(&context, problem)
        if err != nil {
            panic(err)
        } else {
            fmt.Printf("Part 2: %d\n", answer)
        }
    }
}