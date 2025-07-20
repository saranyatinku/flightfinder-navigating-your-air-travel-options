package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"

	"github.com/mateuszmidor/FlightFinder/cmd/finder_cli/cliapp"
)

func main() {
	log.SetPrefix("[APP] ")

	// collect CPU profile
	cpu, _ := os.Create("cpu.out")
	defer cpu.Close()
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()

	// collect traces
	traces, _ := os.Create("trace.out")
	defer traces.Close()
	trace.Start(traces)
	defer trace.Stop()

	flights_data_dir := flag.String("flights_data", "./assets", "-flights_data=./assets")
	flag.Parse()

	cliapp.Run(*flights_data_dir)

	// collect memory profile
	heap, _ := os.Create("mem.out")
	defer heap.Close()
	runtime.GC() // get up-to-date statistics
	pprof.WriteHeapProfile(heap)
}
