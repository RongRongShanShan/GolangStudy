package lib1

func Fun() {
	println("lib1 中的方法 被调用")
}

// init方法在被导包时会被调用
func init() {
	println("lib1 - init 被调用")
}
