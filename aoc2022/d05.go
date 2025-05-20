package aoc2022

import (
	"errors"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var cargoShipLayerRegex *regexp.Regexp
var cargoShipNumbersRegex *regexp.Regexp

func init() {
    cargoShipLayerRegex = regexp.MustCompile(`\[\w\]`)
    cargoShipNumbersRegex = regexp.MustCompile(`^\s*(\d+\s+)*\d*\s*$`)
}

type cargoShip struct {
    towers [][]string
}

func (ship *cargoShip) clone() cargoShip {
    towers := make([][]string, len(ship.towers))
    for i, layer := range ship.towers {
        towers[i] = append(towers[i], layer...)
    }
    return cargoShip{towers}
}

func (ship *cargoShip) move(move *craneMove, crane9001 bool) error {
    if move == nil {
        return errors.New("NOOO")
    }
    split := len(ship.towers[move.src]) - int(move.amount)
    if split < 0 {
        return errors.New("NOOOO")
    }
    remaining, crates := ship.towers[move.src][:split], ship.towers[move.src][split:]
    ship.towers[move.src] = remaining
    if !crane9001 {
        slices.Reverse(crates)
    }
    ship.towers[move.dest] = append(ship.towers[move.dest], crates...)
    return nil
}

func cargoShipGetLayer(s string) []string {
    crateIndices := cargoShipLayerRegex.FindAllIndex([]byte(s), -1)
    layerSize := (len(s) + 1) / 4
    crates := make([]string, layerSize)
    for _, indexPair := range crateIndices {
        start := indexPair[0]
        end := indexPair[1]
        offset := start / 4
        crates[offset] = s[start + 1:end - 1]
    }
    return crates
}

func cargoShipIsLastLine(s string) bool {
    return cargoShipNumbersRegex.Match([]byte(s))
}

var craneMoveRegex *regexp.Regexp

func init() {
    craneMoveRegex = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
}

type craneMove struct {
    amount int64
    src int64
    dest int64
}

func aToCraneMove(s string) (*craneMove, error) {
    matches := craneMoveRegex.FindStringSubmatch(s)
    amount, err := strconv.Atoi(matches[1])
    src, err := strconv.Atoi(matches[2])
    dest, err := strconv.Atoi(matches[3])
    if err != nil {
        return nil, err
    } else {
        return &craneMove{int64(amount), int64(src) - 1, int64(dest) - 1}, nil
    }
}

type Problem05 struct {
    ship cargoShip
    moves []craneMove
}

type Solution05 struct {}

type Day05 struct {}

func (day *Day05) BuildProblem(ctx *Context) (Problem, error) {
    lines := ctx.SplitLines()
    raw_towers := make([][]string, 0)
    i := 0
    ship := cargoShip{nil}
    for ; i < len(lines); i++ {
        if lines[i] == "" {
            i++
            break
        } else if cargoShipIsLastLine(lines[i]) {
            ship.towers = raw_towers
        } else {
            layer := cargoShipGetLayer(lines[i])
            raw_towers = append(raw_towers, layer)
        }
    }
    width := len(raw_towers[len(raw_towers)-1])
    height := len(raw_towers)
    towers := make([][]string, width)
    for x := range width {
        tower := make([]string, 0)
        for y := height - 1; y >= 0; y-- {
            crate := raw_towers[y][x]
            if crate == "" {
                break
            } else {
                tower = append(tower, crate)
            }
        }
        towers[x] = tower
    }
    ship.towers = towers
    moves := make([]craneMove, 0)
    for ; i < len(lines); i++ {
        move, err := aToCraneMove(lines[i])
        if err != nil {
            break
        }
        moves = append(moves, *move)
    }
    return &Problem05{ship, moves}, nil
}

func (day *Day05) BuildSolution(ctx *Context, problem Problem) (Solution, error) {
    return &Solution05{}, nil
}

func (sol *Solution05) solve(p05 *Problem05, crane9001 bool) (string, error) {
    ship := p05.ship.clone()
    for _, move := range p05.moves {
        ship.move(&move, crane9001)
    }
    top := lo.Map(ship.towers, func(tower []string, index int) string {
        return lo.LastOrEmpty(tower)
    })
    return strings.Join(top, ""), nil
}

func (sol *Solution05) P1(ctx *Context, problem Problem) (any, error) {
    p05 := TypeCast[Problem05](problem)
    if p05 == nil {
        return 0, errors.New("Bruh")
    }
    top, err := sol.solve(p05, false)
    if err != nil {
        return 0, err
    } else {
        return top, nil
    }
}

func (sol *Solution05) P2(ctx *Context, problem Problem) (any, error) {
    p05 := TypeCast[Problem05](problem)
    if p05 == nil {
        return 0, errors.New("Bruh")
    }
    top, err := sol.solve(p05, true)
    if err != nil {
        return 0, err
    } else {
        return top, nil
    }
}