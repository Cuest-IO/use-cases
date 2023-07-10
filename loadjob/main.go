package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

var (
	config struct {
		RAMLimit int
		Duration time.Duration
		CPULimit int
	}
	startTime = time.Now()
)

func init() {

	// os.Setenv("CPULimit", "800")
	// os.Setenv("RAMLimit", "10000")
	// os.Setenv("Duration", "10")

	config.CPULimit, _ = strconv.Atoi(os.Getenv("CPULimit"))

	if config.CPULimit > 0 && config.CPULimit <= 1 {
		config.CPULimit = config.CPULimit * 1000
	} else if config.CPULimit > 1000 {
		config.CPULimit = 1000 // 1000Mi, one full CPU
	}

	config.RAMLimit, _ = strconv.Atoi(os.Getenv("RAMLimit"))

	seconds, _ := strconv.Atoi(os.Getenv("Duration"))
	if seconds <= 0 {
		seconds = 60 * 2
	}

	config.Duration = time.Duration(seconds) * time.Second
}

func getRAM(ctx context.Context) int {

	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	return int((memStats.StackInuse + memStats.HeapInuse) / (1 << 20))
}

func loadRAM(ctx context.Context) {

	if config.RAMLimit <= 0 {
		return
	}

	var m1, m2 int
	m1 = getRAM(ctx)
	// 1 Mebibyte = 1048576 bytes
	s := make([]byte, config.RAMLimit*1048576)
	if s != nil {
		m2 = getRAM(ctx)
		fmt.Println("total:", (m2 - m1))

	}
}

func loadCPU(ctx context.Context) {

	onePreiodMillis := 100 * config.CPULimit / 1000
	onePeriodSleep := 100 - onePreiodMillis
	runtime.LockOSThread()
	// endless loop
	for {
		begin := time.Now()
		for {
			// run one period
			if time.Now().Sub(begin).Milliseconds() > int64(onePreiodMillis) {

				// log.Printf("break  time: %d", time.Now().Sub(begin).Milliseconds())
				break
			}
			math.Asinh(float64(math.E*float64(5)) * math.Pow(math.E, float64(6)))

		}
		// sleep
		time.Sleep(time.Duration(onePeriodSleep) * time.Millisecond)
		// log.Printf("run time: %d", time.Now().Sub(begin).Milliseconds())
	}
}

func getCPUCores() int {

	log.Printf("system vCPUs: %d", runtime.NumCPU())
	return runtime.NumCPU()
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), config.Duration)
	defer cancel()

	go loadRAM(ctx)
	go loadCPU(ctx)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM)

	<-ctx.Done()

	if code := os.Getenv("ExitCode"); code != "" {
		i, _ := strconv.Atoi(code)
		os.Exit(i)
	}
}
