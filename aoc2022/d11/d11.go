package aoc2022_d11

import (
	"errors"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/RenoirTan/AdventOfCode2022/aoc2022"
	"github.com/samber/lo"
)

type OperationFunction func (old int64) int64

type Monkey struct {
    Id uint8
    Items []int64
    Operation OperationFunction
    TestInt int64
    TrueDestination uint8
    FalseDestination uint8
}

const MonkeyPattern string = `Monkey ([0-9]+):
\s*Starting items:\s*([0-9]+(\s*,\s*[0-9]+)*)?\s*
\s*Operation: new =\s*(old|[0-9]+)\s*(\*|\+)\s*(old|[0-9]+)\s*
\s*Test: divisible by\s*([0-9]+)\s*
\s*If true: throw to monkey\s*([0-9]+)\s*
\s*If false: throw to monkey\s*([0-9]+)\s*`

var MonkeyRegex *regexp.Regexp = nil

func init() {
    re, err := regexp.Compile(MonkeyPattern)
    if err != nil {
        panic(err)
    } else {
        MonkeyRegex = re
    }
}

func MakeMonkeyFromRaw(
    id string,
    items string,
    op0 string,
    op1 string,
    op2 string,
    test string,
    t string,
    f string,
) (*Monkey, error) {
    parsedId, err := strconv.Atoi(id)
    if err != nil {
        return nil, err
    }
    splitItems := strings.Split(items, ",")
    parsedItems := make([]int64, len(splitItems))
    for i, splitItem := range lo.Map(splitItems, func (item string, index int) string {
        return strings.TrimSpace(item)
    }) {
        parsedItem, err := strconv.Atoi(splitItem)
        if err != nil {
            return nil, err
        }
        parsedItems[i] = int64(parsedItem)
    }
    var binaryOperator func (a int64, b int64) int64 = nil
    switch op1 {
    case "*": binaryOperator = func(a, b int64) int64 { return a * b }
    case "+": binaryOperator = func(a, b int64) int64 { return a + b }
    default:
        return nil, errors.New("Invalid binary operator")
    }
    var parsedOp2 = int64(0)
    if op2 == "old" {
        parsedOp2 = -1
    } else {
        temp, err := strconv.Atoi(op2)
        if err != nil {
            return nil, err
        }
        parsedOp2 = int64(temp)
    }
    var operation OperationFunction = nil
    if parsedOp2 == -1 {
        operation = func (old int64) int64 {
            return binaryOperator(old, old)
        }
    } else {
        operation = func (old int64) int64 {
            return binaryOperator(old, parsedOp2)
        }
    }
    parsedTestInt, err := strconv.Atoi(test)
    if err != nil {
        return nil, err
    }
    parsedTrue, err := strconv.Atoi(t)
    if err != nil {
        return nil, err
    }
    parsedFalse, err := strconv.Atoi(f)
    if err != nil {
        return nil, err
    }
    monkey := &Monkey{
        Id: uint8(parsedId),
        Items: parsedItems,
        Operation: operation,
        TestInt: int64(parsedTestInt),
        TrueDestination: uint8(parsedTrue),
        FalseDestination: uint8(parsedFalse),
    }
    return monkey, nil
}

func MakeMonkeys(lines []string) ([]Monkey, error) {
    fullmatches := MonkeyRegex.FindAllStringSubmatch(strings.Join(lines, "\n"), -1)
    monkeys := make([]Monkey, len(fullmatches))
    for i, fullmatch := range fullmatches {
        id := fullmatch[1]
        items := fullmatch[2]
        op0 := fullmatch[4]
        op1 := fullmatch[5]
        op2 := fullmatch[6]
        test := fullmatch[7]
        t := fullmatch[8]
        f := fullmatch[9]
        monkey, err := MakeMonkeyFromRaw(id, items, op0, op1, op2, test, t, f)
        if err != nil {
            return nil, err
        }
        if monkey == nil {
            return nil, errors.New("Nil monkey returned from MakeMonkeyFromRaw")
        }
        monkeys[i] = *monkey
    }
    return monkeys, nil
}

func (monkey *Monkey) Copy() *Monkey {
    id := monkey.Id
    items := make([]int64, len(monkey.Items))
    copy(items, monkey.Items)
    operation := monkey.Operation
    testInt := monkey.TestInt
    trueDestination := monkey.TrueDestination
    falseDestination := monkey.FalseDestination
    return &Monkey{id, items, operation, testInt, trueDestination, falseDestination}
}

type MonkeyManager struct {
    Monkeys []Monkey
}

func (manager *MonkeyManager) Copy() (*MonkeyManager, error) {
    monkeys := make([]Monkey, len(manager.Monkeys))
    for i, monkey := range manager.Monkeys {
        monkeys[i] = *monkey.Copy()
    }
    return &MonkeyManager{monkeys}, nil
}

type Problem11 struct {
    Manager MonkeyManager
}

type Solution11 struct {}

type Day11 struct {}

func (day *Day11) BuildProblem(ctx *aoc2022.Context) (aoc2022.Problem, error) {
    lines := ctx.SplitLines()
    monkeys, err := MakeMonkeys(lines)
    if err != nil {
        return nil, err
    }
    return &Problem11{MonkeyManager{monkeys}}, nil
}

func (day *Day11) BuildSolution(
    ctx *aoc2022.Context,
    problem aoc2022.Problem,
) (aoc2022.Solution, error) {
    sol := &Solution11{}
    return sol, nil
}

func (monkey *Monkey) RunTurn(reworry func (worry int64) int64) map[uint8][]int64 {
    destinations := map[uint8][]int64 {}
    for _, item := range monkey.Items {
        worry := reworry(monkey.Operation(item))
        var destination uint8 = 0
        if worry % monkey.TestInt == 0 {
            destination = monkey.TrueDestination
        } else {
            destination = monkey.FalseDestination
        }
        destinations[destination] = append(destinations[destination], worry)
    }
    monkey.Items = make([]int64, 0)
    return destinations
}

func (manager *MonkeyManager) RunRound(
    reworry func (worry int64) int64,
    preTurnCallback func (monkey *Monkey),
    postTurnCallback func (monkey *Monkey),
) {
    for i := range len(manager.Monkeys) {
        if preTurnCallback != nil {
            preTurnCallback(&manager.Monkeys[i])
        }
        destinations := manager.Monkeys[i].RunTurn(reworry)
        for j, worries := range destinations {
            manager.Monkeys[j].Items = append(manager.Monkeys[j].Items, worries...)
        }
        if postTurnCallback != nil {
            postTurnCallback(&manager.Monkeys[i])
        }
    }
}

func (sol *Solution11) Solve(
    ctx *aoc2022.Context,
    p11 *Problem11,
    part uint8,
) (any, error) {
    rounds := 20
    reworry := func (worry int64) int64 {
        return worry / 3
    }
    if part == 2 {
        // https://todd.ginsberg.com/post/advent-of-code/2022/day11/
        rounds = 10000
        residueRange := lo.Reduce(
            lo.Map(p11.Manager.Monkeys, func (item Monkey, index int) int64 {
                return item.TestInt
            }),
            func (agg int64, item int64, index int) int64 {
                return agg * item
            },
            1,
        )
        reworry = func (worry int64) int64 {
            return worry % residueRange
        }
    }

    inspectionLeaderboard := map[uint8]uint64 {}
    manager, err := p11.Manager.Copy()
    if err != nil {
        return nil, err
    }
    for _ = range rounds {
        manager.RunRound(reworry, func (monkey *Monkey) {
            id := monkey.Id
            score := uint64(len(monkey.Items))
            inspectionLeaderboard[id] += score
        }, nil)
    }
    sortedScores := lo.Values(inspectionLeaderboard)
    slices.Sort(sortedScores)
    // fmt.Println(inspectionLeaderboard)
    n := len(sortedScores)
    monkeyBusiness := int64(sortedScores[n - 2]) * int64(sortedScores[n - 1])
    return monkeyBusiness, nil
}

func (sol *Solution11) P1(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p11 := aoc2022.TypeCast[Problem11](problem)
    if p11 == nil {
        return nil, errors.New("Bruh")
    }
    return sol.Solve(ctx, p11, 1)
}

func (sol *Solution11) P2(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p11 := aoc2022.TypeCast[Problem11](problem)
    if p11 == nil {
        return nil, errors.New("Bruh")
    }
    return sol.Solve(ctx, p11, 2)
}