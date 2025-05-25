package aoc2022_d07

import (
	"errors"
	"fmt"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/RenoirTan/AdventOfCode2022/aoc2022"
	"github.com/samber/lo"
)

type Problem07 struct {
    pathsToSize map[string]int64
}

type Solution07 struct {
    dirsSizes map[string]int64
}

type Day07 struct {}

func (day *Day07) BuildProblem(ctx *aoc2022.Context) (aoc2022.Problem, error) {
    pathsToSize := map[string]int64{"/": -1}
    currentPath := "/"
    for _, line := range ctx.SplitLines() {
        if strings.HasPrefix(line, "$ cd ") {
            currentPath = path.Join(currentPath, line[5:])
            pathsToSize[currentPath] = -1
        } else if strings.HasPrefix(line, "$ ls") {
            continue
        } else if strings.HasPrefix(line, "dir ") {
            nodePath := path.Join(currentPath, line[5:])
            pathsToSize[nodePath] = -1
        } else {
            components := strings.SplitN(line, " ", 2)
            if len(components) != 2 {
                return nil, errors.New(fmt.Sprintf("Could not interpret %s as size and file", line))
            }
            size, err := strconv.Atoi(components[0])
            if err != nil {
                return nil, err
            }
            nodePath := path.Join(currentPath, components[1])
            pathsToSize[nodePath] = int64(size)
        }
    }
    return &Problem07{pathsToSize}, nil
}

func (day *Day07) BuildSolution(
    ctx *aoc2022.Context,
    problem aoc2022.Problem,
) (aoc2022.Solution, error) {
    return &Solution07{nil}, nil
}

func isDirectDescendantOf(parent, child string) bool {
    parent = path.Clean(parent)
    child = path.Clean(child)
    poc := path.Join(child, "..")
    return poc != child && parent == poc
}

func (p07 *Problem07) directDescendantsOf(parent string) []string {
    return lo.FilterMapToSlice(p07.pathsToSize, func(p string, s int64) (string, bool) {
        return p, isDirectDescendantOf(parent, p)
    })
}

func (p07 *Problem07) sizeOf(parent string) int64 {
    size := p07.pathsToSize[parent]
    if size > 0 {
        return size
    } else {
        return lo.SumBy(p07.directDescendantsOf(parent), func(child string) int64 {
            return p07.sizeOf(child)
        })
    }
}

func (p07 *Problem07) sizeOfDirs() map[string]int64 {
    dirs := lo.FilterMapToSlice(p07.pathsToSize, func(p string, s int64) (string, bool) {
        return p, s <= 0
    })
    sizes := lo.SliceToMap(dirs, func(p string) (string, int64) {
        return p, p07.sizeOf(p)
    })
    return sizes
}

func (sol *Solution07) P1(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p07 := aoc2022.TypeCast[Problem07](problem)
    if p07 == nil {
        return 0, errors.New("Bruh")
    }
    dirsSizes := p07.sizeOfDirs()
    sol.dirsSizes = dirsSizes
    return lo.SumBy(lo.Values(dirsSizes), func(s int64) int64 {
        if s <= 100000 { return s } else { return 0 }
    }), nil
}

func (sol *Solution07) P2(ctx *aoc2022.Context, problem aoc2022.Problem) (any, error) {
    p07 := aoc2022.TypeCast[Problem07](problem)
    if p07 == nil {
        return 0, errors.New("Bruh")
    }
    if sol.dirsSizes == nil {
        sol.dirsSizes = p07.sizeOfDirs()
    }
    used := sol.dirsSizes["/"]
    neededSpace := 30000000 - (70000000 - used)
    dirs := lo.Keys(sol.dirsSizes)
    sort.Slice(dirs, func (i, j int) bool {
        a := dirs[i]
        b := dirs[j]
        return sol.dirsSizes[a] < sol.dirsSizes[b]
    })
    for _, p := range dirs {
        size := sol.dirsSizes[p]
        if size >= neededSpace {
            return size, nil
        }
    }
    return nil, errors.New("Could not find answer")
}