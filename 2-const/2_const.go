package main

import "unsafe"

/**
常量的定义
*/

func main() {
	// 常量定义格式  const identifier [type] = value
	const length = 5
	// 多重赋值
	const a, b, c = 1, "a", 0.1

	// 常量用于枚举
	const (
		aaa = "abc"
		bbb = len(aaa) // 内置函数作常量表达式(常量表达式中，函数必须是内置函数)
		ccc = unsafe.Sizeof(aaa)
	)
	println(aaa, bbb, ccc) // 输出结果：abc 3 16

	// iota 自增长
	const (
		ia = iota // 0
		ib        // 1
		ic        // 2
	)
	println(ia, ib, ic) // 输出结果：0 1 2

	// iota与表达式
	const (
		ai = 2 * iota       // iota=0, ai=0  iota照样从0增长，变量值为表达式计算结果
		bi                  // iota=1, bi=2  没有表达式时按之前的表达式进行运算
		ci                  // iota=2, bi=4
		di = (iota + 1) / 2 // iota=3, di=2  更换表达时后按新表达式进行运算，iota正常增长
		ei                  // iota=4, di=2
		fi                  // iota=5, di=3
	)
	println(ai, bi, ci, di, ei, fi) // 输出结果：0 2 4 2 2 3

	// 两个常量定义在一行
	const (
		_, _   = iota + 1, iota + 2 // _变量无法访问，相当于iota=0的情况被忽略
		ca, cb                      // iota=0, ca=2, cb=3
		cc, cd                      // iota=1, cc=3, cb=4
		ce, cf                      // ..
	)
	// 无法访问空白标识符_
	//println(_)
	println(ca, cb, cc, cd, ce, cf) // 输出结果：2 3 3 4 4 5
}
