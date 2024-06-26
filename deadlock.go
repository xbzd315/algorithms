package main

import (
    "fmt"
    "runtime"
	_"time"
    "sync"
)

// 定义两把锁
var (
    lockA = &sync.Mutex{}
    lockB = &sync.Mutex{}
)

// 资源互斥条件
// 持有并等待条件
// 不可剥夺条件
// 环路等待条件

func threadA() {
    lockA.Lock()
    fmt.Println("Lock A acquired by Thread A")
    // time.Sleep(1000 * time.Millisecond)
    // 尝试获取锁B，但不会释放锁A
    lockB.Lock()
    defer lockB.Unlock()
    fmt.Println("Lock B acquired by Thread A")

    // 正常执行完毕，释放锁A
    lockA.Unlock()
}

func threadB() {
    // time.Sleep(1000 * time.Millisecond)
    lockB.Lock()
    fmt.Println("Lock B acquired by Thread B")
    // time.Sleep(1000 * time.Millisecond)
    // 尝试获取锁A，但不会释放锁B
    lockA.Lock()
    defer lockA.Unlock()
    fmt.Println("Lock A acquired by Thread B")

    // 正常执行完毕，释放锁B
    lockB.Unlock()
}

func main() {
    // go threadA()
    // go threadB()
    // 让出时间让goroutine有机会执行
    // 实际情况下，死锁是由于两个线程在无法继续执行时等待对方释放锁
    // 因此，这里通过sleep让出足够的时间来发生死锁

    // 分配一个逻辑处理器给调度器使用
    runtime.GOMAXPROCS(2)
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        defer wg.Done()
        threadA()
    }()
    go func() {
        defer wg.Done()
        threadB()
    }()
    wg.Wait()
}