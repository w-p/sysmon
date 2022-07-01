package pkg

import (
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// Stats represents system stats
type Stats map[string]float64

// GenerateStats produces system stats
func GenerateStats(t time.Duration) Stats {
	stats := Stats{
		"cpu":  GetCPU(t),
		"ram":  GetRAM(),
		"ping": GetPing("8.8.8.8"),
	}

	return stats
}

// GetCPU gets cpu usage
func GetCPU(t time.Duration) float64 {
	c, err := cpu.Percent(t, false)
	if err != nil {
		panic("failed to get cpu stats")
	}

	return c[0]
}

// GetRAM gets memory usage
func GetRAM() float64 {
	m, err := mem.VirtualMemory()
	if err != nil {
		panic("failed to get memory stats")
	}

	return m.UsedPercent
}

// GetPing gets ping time to a given address
func GetPing(address string) float64 {
	bytes, err := exec.Command("ping", address, "-c 1").Output()
	if err != nil {
		// panic("failed to get ping time")
		return 0.0
	}

	output := string(bytes)
	line := strings.Split(output, "\n")[1]
	latency := strings.Split(line, " ")[6]
	ms := strings.Split(latency, "=")[1]
	result, err := strconv.ParseFloat(ms, 64)

	if err != nil {
		panic("failed to parse ping output")
	}

	return result
}
