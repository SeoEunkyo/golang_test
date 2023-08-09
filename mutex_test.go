package playGround

import (
	"fmt"
	"sync"
	"testing"
)

type LockMap struct {
	Lock sync.RWMutex
	Data map[string][]int
}

func CreateLockMap() *LockMap {
	return &LockMap{
		Lock: sync.RWMutex{},
		Data: make(map[string][]int),
	}

}

func (l *LockMap) ReadMap(key string) []int {
	l.Lock.RLock()
	defer l.Lock.RUnlock()
	val, ok := l.Data[key]
	if !ok {
		return nil
	}
	return val
}

func (l *LockMap) AddKey(key string) {
	l.Lock.Lock()
	defer l.Lock.Unlock()
	ints := []int{1, 2, 3}
	fmt.Printf("address : %p \n", &ints)
	l.Data[key] = ints
	val := l.Data[key]
	fmt.Printf("map address : %p \n", &val)

}

func (l *LockMap) Print(key string) {
	ints := l.ReadMap(key)
	fmt.Println("ints", ints)
}

func TestMutex(t *testing.T) {
	testData := CreateLockMap()
	testData.AddKey("a")

	targets := testData.ReadMap("a")
	fmt.Printf("address : %p \n", &targets)

	for i, target := range targets {
		targets[i] = target + 10
	}
	fmt.Printf("address : %p \n", &targets)
	testData.Print("a")
}

func TestArrayCopy(t *testing.T) {
	ints := []int{1, 2, 3}
	pointer := &ints
	fmt.Printf("address : %p \n", ints)
	fmt.Printf("address : %p \n", pointer)
	match := *pointer
	match[0] = 10
	fmt.Printf("values : %v \n", ints)
	fmt.Printf("values : %v \n", pointer)
}
