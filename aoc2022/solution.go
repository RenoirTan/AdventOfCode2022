package aoc2022

type Solution interface{
    P1(*Context, Problem) (any, error)
    P2(*Context, Problem) (any, error)
}