package dxgi

import (
	"testing"
)

func TestDXGIKeyedMutex_AcquireSync(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 键控互斥锁需要从资源创建
	t.Log("KeyedMutex tests require resource object")
}

func TestDXGIKeyedMutex_ReleaseSync(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	// 需要实际的键控互斥锁对象
	t.Log("KeyedMutex tests require actual mutex object")
}
