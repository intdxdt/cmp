package cmp

import "github.com/intdxdt/math"

type Compare func(a, b interface{}) int

func Int(a, b interface{}) int {
    d := a.(int) - b.(int)
    if d > 0 {
        return 1
    } else if d < 0 {
        return -1
    }
    return 0
}

func Str(a, b interface{}) int {
    var as = a.(string)
    var bs = b.(string)
    if as > bs {
        return 1
    } else if as < bs {
        return -1
    }
    return 0
}

func F64(a, b interface{}) int {
    d := a.(float64) - b.(float64)
    if math.FloatEqual(d, 0) {
        return 0
    } else if d < 0 {
        return -1
    }
    return 1
}
