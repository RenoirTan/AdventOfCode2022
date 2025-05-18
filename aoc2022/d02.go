package aoc2022

import (
	"errors"
)

type hand int64
const (
    rock hand = 1
    paper hand = 2
    scissors hand = 3
)

type result int64
const (
    lose result = 0
    draw result = 3
    win result = 6
)

func symToHand(sym string) hand {
    switch sym {
    case "A": return rock
    case "X": return rock
    case "B": return paper
    case "Y": return paper
    case "C": return scissors
    case "Z": return scissors
    default: panic("what")
    }
}

func gameResult(opponent hand, self hand) result {
    if opponent == self {
        return draw
    } else if (self == rock && opponent == scissors) ||
        (self == paper && opponent == rock) ||
        (self == scissors && opponent == paper) {
        return win
    } else {
        return lose
    }
}

func handThatWinsAgainst(h hand) hand {
    switch h {
    case rock: return paper
    case paper: return scissors
    case scissors: return rock
    default: panic("what")
    }
}

func handThatLosesTo(h hand) hand {
    switch h {
    case rock: return scissors
    case paper: return rock
    case scissors: return paper
    default: panic("what")
    }
}

type Problem02 struct {
    Plays [][2]string
}

type Solution02 struct{}

type Day02 struct{}

func (day *Day02) BuildProblem(ctx *Context) (Problem, error) {
    lines := ctx.SplitLines()
    n_plays := len(lines)
    plays := make([][2]string, n_plays)
    for i, line := range lines {
        plays[i][0] = line[0:1]
        plays[i][1] = line[2:3]
    }
    return &Problem02{plays}, nil
}

func (day *Day02) BuildSolution(ctx *Context,problem Problem) (Solution, error) {
    return &Solution02{}, nil
}

func (sol *Solution02) P1(ctx *Context, problem Problem) (int64, error) {
    p02 := TypeCast[Problem02](problem)
    if p02 == nil {
        return 0, errors.New("bruh")
    }
    var score int64 = 0
    for _, play := range p02.Plays {
        opponent := symToHand(play[0])
        self := symToHand(play[1])
        resultScore := gameResult(opponent, self)
        score += int64(resultScore) + int64(self)
    }
    return score, nil
}

func (sol *Solution02) P2(ctx *Context, problem Problem) (int64, error) {
    p02 := TypeCast[Problem02](problem)
    if p02 == nil {
        return 0, errors.New("bruh")
    }
    var score int64 = 0
    for _, play := range p02.Plays {
        opponent := symToHand(play[0])
        switch play[1] {
        case "X":
            score += int64(lose) + int64(handThatLosesTo(opponent))
        case "Y":
            score += int64(draw) + int64(opponent)
        case "Z":
            score += int64(win) + int64(handThatWinsAgainst(opponent))
        }
    }
    return score, nil
}