package aoc2022

type Int64RangeInclusive struct {
    Lower int64
    Upper int64
}

func (r *Int64RangeInclusive) Contains(i int64) bool {
    return r.Lower <= i && i <= r.Upper
}

func (r *Int64RangeInclusive) Intersection(s *Int64RangeInclusive) *Int64RangeInclusive {
    if r == nil || s == nil {
        return nil
    }
    lower := max(r.Lower, s.Lower)
    upper := min(r.Upper, s.Upper)
    if lower > upper {
        return nil
    } else {
        return &Int64RangeInclusive{lower, upper}
    }
}

func (r *Int64RangeInclusive) IsSubsetOf(s *Int64RangeInclusive) bool {
    intersection := r.Intersection(s)
    if intersection == nil {
        return false
    } else {
        return *intersection == *r
    }
}