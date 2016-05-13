package one

import (
    "sort"
)

type CharCount struct {
    mapping map[string]int
    order []string
}

func (cc *CharCount) Len() int {
    return len(cc.mapping)
}

func (cc *CharCount) Swap(i, j int) {
    cc.order[i], cc.order[j] = cc.order[j], cc.order[i]
}

func (cc *CharCount) Less(i, j int) bool {
    return cc.mapping[cc.order[i]] > cc.mapping[cc.order[j]]
}

func MakeCharCount(s string) *CharCount {
    out := new(CharCount)

    out.mapping = make(map[string]int)
    for i := 0; i < len(s); i += 2 {
        out.mapping[s[i:i+2]] += 1
    }
    n := len(out.mapping)
    out.order = make([]string, n)

    i := 0
    for k := range out.mapping {
        out.order[i] = k
        i++
    }
    sort.Sort(out)
    return out
}
