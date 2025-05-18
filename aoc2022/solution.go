package aoc2022

type Solution interface{
    P1(*Context, Problem) (int64, error)
    P2(*Context, Problem) (int64, error)
}