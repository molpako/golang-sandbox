package multi

import (
	"sync"
	"syscall"
)

func slowSyscall() {
	fd, _ := syscall.Socket(
		syscall.AF_INET,
		syscall.SOCK_DGRAM,
		syscall.IPPROTO_UDP,
	)
	// 0.1秒かかるようにする
	timeout := &syscall.Timeval{Sec: 0, Usec: 100000}
	syscall.Select(fd, nil, nil, nil, timeout)
}

func Successive() int {
	// 逐次的に関数を実行させる
	for i := 0; i < 10; i++ {
		slowSyscall()
	}
	return 0
}

// --- PASS: TestSuccessive/#00 (1.02s)

func Concurrency() int {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// goroutineを立ち上げ関数を実行させる
		go func() {
			slowSyscall()
			wg.Done()
		}()
	}
	// 複数のgoroutineの処理が終わるまで待機する
	wg.Wait()
	return 0
}

// --- PASS: TestCunccurency/#00 (0.10s)
