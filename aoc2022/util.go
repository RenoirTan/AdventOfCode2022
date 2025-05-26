package aoc2022

import (
	"cmp"
	"regexp"
	"sort"
)

var splitLinesPattern *regexp.Regexp

func init() {
    splitLinesPattern = regexp.MustCompile("(\r\n|\r|\n)")
}

func SplitLines(s string) []string {
    return splitLinesPattern.Split(s, -1)
}

func TypeCast[T any](object any) *T {
    switch object.(type) {
    case T:
        return object.(*T)
    case *T:
        return object.(*T)
    default:
        return nil
    }
}

func Sort[T cmp.Ordered](s []T) {
    sort.Slice(s, func(i, j int) bool {
        return s[i] < s[j]
    })
}

func BToi(b bool) int {
    if b {
        return 1
    } else {
        return 0
    }
}