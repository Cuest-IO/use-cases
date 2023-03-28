package main

import (
	"context"
	"log"
	"math"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	config struct {
		RAM      int
		Duration time.Duration
		CPUSleep int64
	}
	startTime = time.Now()
)

func init() {

	config.CPUSleep, _ = strconv.ParseInt(os.Getenv("CPUSleep"), 10, 64)
	config.RAM, _ = strconv.Atoi(os.Getenv("RAM"))

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

	if config.RAM <= 0 {
		return
	}

	var (
		data []string
	)

	for {

		if getRAM(ctx) >= config.RAM {
			time.Sleep(time.Second * 10)
		} else {
			data = append(data, strings.Repeat("a", 1000))
			time.Sleep(time.Microsecond)
		}
	}

}

func compute(ctx context.Context) {
	for i := 0; i < 50; i++ {
		go func(i int) {
			for {
				math.Asinh(float64(math.E*float64(i)) * math.Pow(math.E, float64(i)))
				if config.CPUSleep > 0 {
					time.Sleep(time.Nanosecond * time.Duration(config.CPUSleep))
				}
			}
		}(i)
	}
}

func loadCPU(ctx context.Context) {

	go compute(ctx)
	for {

		cores, err := getCPUCores()
		if err != nil {

			log.Println(err)
			time.Sleep(time.Second * 10)
			continue

		}

		mem := getRAM(ctx)
		log.Printf("cores: %0.4f; Mem %d mb\n", cores, mem)
		time.Sleep(time.Second)
	}
}

func getCPUCores() (float64, error) {

	var rusage syscall.Rusage
	if err := syscall.Getrusage(0, &rusage); err != nil {
		return 0, err
	}

	gcstat := new(debug.GCStats)
	debug.ReadGCStats(gcstat)
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)

	return (float64(rusage.Stime.Nano()+rusage.Utime.Nano()) / float64(time.Now().UnixNano()-startTime.UnixNano())), nil
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
