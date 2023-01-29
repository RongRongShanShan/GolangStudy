package main

import "testing"

//测试文件  以_test.go结尾

// 单元测试函数命名格式  func TestXxx(*testing.T)
func TestManyGoWait(t *testing.T) {
	ManyGoWait()
}

func TestHelloGoRoutine(t *testing.T) {
	HelloGoRoutine()
}

// 基准测试函数命名格式  func BenchmarkXxx(*testing.B)
func BenchmarkManyGoWait(b *testing.B) {
	ManyGoWait()
}
