package aoc2022_d10

import (
	"errors"
	"strconv"
	"strings"

	"github.com/RenoirTan/AdventOfCode2022/aoc2022"
	"github.com/samber/lo"
)

type Mnemonic uint8

const (
    NoOp Mnemonic = iota
    AddX
)

func (m Mnemonic) InstructionDuration() uint64 {
    switch m {
    case NoOp:
        return 1
    case AddX:
        return 2
    default:
        return 0
    }
}

type Instruction struct {
    Mnemonic Mnemonic
    Value int64
}

func MakeInstruction(s string) (*Instruction, error) {
    if (len(s) < 4) {
        return nil, errors.New("Bad instruction")
    }
    instruction := &Instruction{}
    switch s[:4] {
    case "noop":
        instruction.Mnemonic = NoOp
    case "addx":
        instruction.Mnemonic = AddX
    default:
        return nil, errors.New("Invalid mnemonic")
    }
    switch instruction.Mnemonic {
    case AddX:
        i, err := strconv.Atoi(s[5:])
        if err != nil {
            return nil, err
        }
        instruction.Value = int64(i)
    }
    return instruction, nil
}

type Problem10 struct {
    Instructions []Instruction
}

type ProcessorState struct {
    Instruction Instruction
    RemainingCycles uint64
    TotalCycles uint64
    X int64
}

func InitProcessorState() ProcessorState {
    return ProcessorState{Instruction{NoOp, 0}, 0, 1, 1}
}

// Returns true if nextInstruction is consumed
func (state *ProcessorState) evolve(nextInstruction *Instruction) bool {
    // fmt.Println(state)
    // fmt.Println(nextInstruction)
    // Load new instruction if required
    consumed := false
    if state.RemainingCycles == 0 {
        if nextInstruction != nil {
            state.Instruction = *nextInstruction
            state.RemainingCycles = state.Instruction.Mnemonic.InstructionDuration()
        } else {
            return true
        }
        consumed = true
    }

    switch state.Instruction.Mnemonic {
    case NoOp:
        break
    case AddX:
        if state.RemainingCycles == 1 {
            state.X += state.Instruction.Value
        }
    }
    state.RemainingCycles--
    state.TotalCycles++

    return consumed
}

type Solution10 struct {}

type Day10 struct {}

func (day *Day10) BuildProblem(ctx *aoc2022.Context) (aoc2022.Problem, error) {
    lines := ctx.SplitLines()
    instructions := make([]Instruction, len(lines))
    for i, line := range lines {
        instruction, err := MakeInstruction(line)
        if err != nil {
            return nil, err
        }
        instructions[i] = *instruction
    }
    return &Problem10{instructions}, nil
}

func (day *Day10) BuildSolution(
    ctx *aoc2022.Context,
    problem aoc2022.Problem,
) (aoc2022.Solution, error) {
    sol := &Solution10{}
    return sol, nil
}

func (sol *Solution10) Solve(
    ctx *aoc2022.Context,
    p10 *Problem10,
    preCallback func (i int, state ProcessorState),
    postCallback func (i int, state ProcessorState),
) error {
    state := InitProcessorState()
    programLength := len(p10.Instructions)
    for i := 0; i <= programLength; {
        if preCallback != nil {
            preCallback(i, state)
        }
        // fmt.Println(i)
        var nextInstruction *Instruction = nil
        if i < programLength {
            nextInstruction = &p10.Instructions[i]
        }
        consumed := state.evolve(nextInstruction)
        if consumed {
            i++
        }
        if postCallback != nil {
            postCallback(i, state)
        }
    }
    return nil
}

func (sol *Solution10) P1(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p10 := aoc2022.TypeCast[Problem10](problem)
    if p10 == nil {
        return nil, errors.New("Bruh")
    }
    totalSignalStrength := int64(0)
    interestingCycles := []uint64{20, 60, 100, 140, 180, 220}
    err := sol.Solve(ctx, p10, nil, func(i int, state ProcessorState) {
        if lo.Contains(interestingCycles, state.TotalCycles) {
            totalSignalStrength += state.X * int64(state.TotalCycles)
        }
    })
    if err != nil {
        return nil, err
    } else {
        return totalSignalStrength, nil
    }
}

func (sol *Solution10) P2(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p10 := aoc2022.TypeCast[Problem10](problem)
    if p10 == nil {
        return nil, errors.New("Bruh")
    }
    crtState := [240]bool{}
    err := sol.Solve(ctx, p10, func(i int, state ProcessorState) {
        phase := state.TotalCycles - 1
        rowPhase := int64(phase % 40)
        if state.X - 1 <= rowPhase && rowPhase <= state.X + 1 {
            crtState[phase] = true
        }
    }, nil)
    if err != nil {
        return nil, err
    } else {
        result := ""
        for i := 0; i < 240; i += 40 {
            row := strings.Join(lo.Map(crtState[i:i+40], func (item bool, index int) string {
                if item {
                    return "#"
                } else {
                    return "."
                }
            }), "")
            result += "\n" + row
        }
        return result, nil
    }
}