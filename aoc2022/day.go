package aoc2022

type Day interface{
    BuildProblem(*Context) (Problem, error)
    BuildSolution(*Context, Problem) (Solution, error)
}