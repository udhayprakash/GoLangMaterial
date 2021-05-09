## cloning slice

    func MakeCopy(s []int) []int {10%
        r := make([]int, len(s))
        copy(r, s)
        return r
    }

    func MakeAppend(s []int) []int {
        return append(make([]int, 0, len(s)), s...)
    }

    func MakeAppend2(s []int) []int {
        r := make([]int, 0, len(s))
        return append(r, s...)
    }

    It looks MakeCopy is always (about 10%) faster than MakeAppend ones.

    The MakeAppend ways are faster than the following append ways for small slices,
    but slower than them for large slices, which is also some surprising.

    func Append(s []int) []int {
        return append([]int(nil), s...)
    }

    func Append2(s []int) []int {
        return append(s[:0:0], s...)
    }
