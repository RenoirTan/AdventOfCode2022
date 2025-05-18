package aoc2022

type Day interface{
    BuildProblem(*Context) (Problem, error)
    BuildSolution(*Context, Problem) (Solution, error)
}

func GetDay(d uint64) Day {
    return []Day{
        nil,
        &Day01{},
        &Day02{},
    }[d]
}