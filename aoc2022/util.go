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

type Int interface { int | int8 | int16 | int32 | int64 }

type UInt interface { uint | uint8 | uint16 | uint32 | uint64 }

type Float interface { float32 | float64 }

type Complex interface { complex64 | complex128 }

type Scalar interface { Int | UInt | Float }

type Number interface { Int | UInt | Float | Complex }

func AbsScalar[T Scalar](scalar T) T {
    return max(scalar, -scalar)
}

func NormalizeScalar[T Scalar](scalar T) T {
    if scalar == 0 {
        return scalar
    } else {
        return scalar / AbsScalar(scalar)
    }
}