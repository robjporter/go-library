package xping

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/robjporter/go-library/xping/ping"
)

type MacPinger struct{}

func New() Pinger {
	return &MacPinger{}
}

func (p *MacPinger) Ping(hostname string, times int, delay int) Result {
	result := Result{}
	result.Latency.Min = 1000.0

	var latencySum float64

	errors := 0
	for i := 0; i < times; i++ {
		now := time.Now()
		err := ping.Pinger(hostname, 1)
		end := time.Now()

		latencyValue := end.Sub(now).Seconds() * 1000.0
		latencySum += latencyValue
		if latencyValue > result.Latency.Max {
			result.Latency.Max = latencyValue
		}
		if latencyValue < result.Latency.Min {
			result.Latency.Min = latencyValue
		}
		if err != nil {
			errors++
			continue
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}
	result.PacketLoss = int64(errors) * 100 / int64(times)
	result.Latency.Avg = latencySum / float64(times)
	result.Online = true
	if result.PacketLoss == 100 {
		result.Online = false
	}
	return result
}

type MacFallbackPinger struct {
}

func NewMacFallbackPinger() *MacFallbackPinger {
	return &MacFallbackPinger{}
}

func (p *MacFallbackPinger) Ping(hostname string, times int, delay int) Result {
	cmd := exec.Command("/sbin/ping", hostname, "-c", fmt.Sprintf("%d", times), "-i", fmt.Sprintf("%d", delay))
	output, _ := cmd.CombinedOutput()
	outputStr := string(output)

	result := Result{}
	result.Online = !strings.Contains(outputStr, "100% packet loss")

	if result.Online {
		outputLines := strings.Split(outputStr, "\n")
		for _, line := range outputLines {
			if strings.Contains(line, "packet loss") {
				markers := strings.Split(line, ", ")
				markers = strings.Split(markers[2], "%")
				result.PacketLoss = giveInt64(markers[0])
			}
			if strings.Contains(line, "round-trip") {
				markers := strings.Split(line, " = ")
				markers = strings.Split(markers[1], "/")
				result.Latency = Latency{
					giveFloat(markers[0]),
					giveFloat(markers[1]),
					giveFloat(markers[2]),
				}
				break
			}
		}
	}
	return result
}
