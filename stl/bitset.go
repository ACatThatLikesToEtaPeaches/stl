package stl

import "math/bits"

// Bitset
// 参考 C++ 的标准库源码 https://gcc.gnu.org/onlinedocs/libstdc++/libstdc++-html-USERS-3.4/bitset-source.html
const _w = bits.UintSize

func NewBitset(n int) Bitset { return make(Bitset, n/_w+1) } // (n+_w-1)/_w

type BitsetInterface interface {
	Has(p int) bool     // get bitset[p]
	Flip(p int)         // 翻转
	Set(p int)          // 置 1
	Reset(p int)        // 置 0
	OnesCount() (c int) // 返回 1 的个数
	All0(l, r int) bool // 判断 [l,r] 范围内的数是否全为 0
	All1(l, r int) bool // 判断 [l,r] 范围内的数是否全为 1
}

type Bitset []uint

func (b Bitset) Has(p int) bool { return b[p/_w]&(1<<(p%_w)) != 0 } // get
func (b Bitset) Flip(p int)     { b[p/_w] ^= 1 << (p % _w) }        // 翻转
func (b Bitset) Set(p int)      { b[p/_w] |= 1 << (p % _w) }        // 置 1
func (b Bitset) Reset(p int)    { b[p/_w] &^= 1 << (p % _w) }       // 置 0

// 左移 k 位
// 应用 https://leetcode-cn.com/problems/minimize-the-difference-between-target-and-chosen-elements/submissions/
func (b Bitset) Lsh(k int) {
	if k == 0 {
		return
	}
	shift, offset := k/_w, k%_w
	if shift >= len(b) {
		for i := range b {
			b[i] = 0
		}
		return
	}
	if offset == 0 {
		// Fast path
		copy(b[shift:], b)
	} else {
		for i := len(b) - 1; i > shift; i-- {
			b[i] = b[i-shift]<<offset | b[i-shift-1]>>(_w-offset)
		}
		b[shift] = b[0] << offset
	}
	for i := 0; i < shift; i++ {
		b[i] = 0
	}
}

// 右移 k 位
func (b Bitset) Rsh(k int) {
	if k == 0 {
		return
	}
	shift, offset := k/_w, k%_w
	if shift >= len(b) {
		for i := range b {
			b[i] = 0
		}
		return
	}
	lim := len(b) - 1 - shift
	if offset == 0 {
		// Fast path
		copy(b, b[shift:])
	} else {
		for i := 0; i < lim; i++ {
			b[i] = b[i+shift]>>offset | b[i+shift+1]<<(_w-offset)
		}
		// 注意：若前后调用 lsh 和 rsh，需要注意超出 n 的范围的 1 对结果的影响
		b[lim] = b[len(b)-1] >> offset
	}
	for i := lim + 1; i < len(b); i++ {
		b[i] = 0
	}
}

// 返回 1 的个数
func (b Bitset) OnesCount() (c int) {
	for _, v := range b {
		c += bits.OnesCount(v)
	}
	return
}

// 遍历所有 1 的位置
// 如果对范围有要求，可在 f 中 return p < n
func (b Bitset) Foreach(f func(p int) (Break bool)) {
	for i, v := range b {
		for ; v > 0; v &= v - 1 {
			j := i*_w | bits.TrailingZeros(v)
			if f(j) {
				return
			}
		}
	}
}

// 返回第一个 0 的下标，不存在时会返回一个不小于 n 的位置
func (b Bitset) Index0() int {
	for i, v := range b {
		if ^v != 0 {
			return i*_w | bits.TrailingZeros(^v)
		}
	}
	return len(b) * _w
}

// 返回第一个 1 的下标，不存在时会返回一个不小于 n 的位置（同 C++ 中的 _Find_first）
func (b Bitset) Index1() int {
	for i, v := range b {
		if v != 0 {
			return i*_w | bits.TrailingZeros(v)
		}
	}
	return len(b) * _w
}

// 返回下标严格大于 p 的第一个 1 的下标，不存在时会返回一个不小于 n 的位置（同 C++ 中的 _Find_next）
func (b Bitset) Next1(p int) int {
	p++ // make bound inclusive
	if i := p / _w; i < len(b) {
		v := b[i] & (^uint(0) << (p % _w)) // mask off bits below bound
		if v != 0 {
			return i*_w | bits.TrailingZeros(v)
		}
		for i++; i < len(b); i++ {
			if b[i] != 0 {
				return i*_w | bits.TrailingZeros(b[i])
			}
		}
	}
	return len(b) * _w
}

// 判断 [l,r] 范围内的数是否全为 0
// https://codeforces.com/contest/1107/problem/D（标准做法是二维前缀和）
func (b Bitset) All0(l, r int) bool {
	i := l / _w
	if i == r/_w {
		return b[i]>>(l%_w)&(1<<(r-l+1)-1) == 0
	}
	if b[i]>>(l%_w) != 0 {
		return false
	}
	for i++; i < r/_w; i++ {
		if b[i] != 0 {
			return false
		}
	}
	return b[r/_w]&(1<<(r%_w+1)-1) == 0
}

// 判断 [l,r] 范围内的数是否全为 1
func (b Bitset) All1(l, r int) bool {
	i := l / _w
	if i == r/_w {
		mask := uint(1)<<(r-l+1) - 1
		return b[i]>>(l%_w)&mask == mask
	}
	if ^(b[i] | (1<<(l%_w) - 1)) != 0 {
		return false
	}
	for i++; i < r/_w; i++ {
		if ^b[i] != 0 {
			return false
		}
	}
	mask := uint(1)<<(r%_w+1) - 1
	return b[r/_w]&mask == mask
}
