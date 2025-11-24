package dbg

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func ProfileTime() {
	defer func(start time.Time) {
		log.Printf("took %v\n", time.Since(start))
	}(time.Now())
}

// TODO
func ProfileCPU() {
	cpu, err := os.Create("cpu.pprof")

	if err != nil {
		log.Fatal(err)
	}

	defer cpu.Close()

	defer pprof.StopCPUProfile()

	if err := pprof.StartCPUProfile(cpu); err != nil {
		log.Fatal(err)
	}
}

// TODO
func ProfileMem() {
	mem, err := os.Create("mem.pprof")

	if err != nil {
		log.Fatal(err)
	}

	defer mem.Close()

	if err := pprof.WriteHeapProfile(mem); err != nil {
		log.Fatal(err)
	}
}
