package pomodoro

import (
	"fmt"
	"time"
)

const tickStep = 1 * time.Second
const workDuration = 20 * time.Minute
const restDuration = 5 * time.Minute

type PomodoroStatus int

const (
	_ PomodoroStatus = iota
	Stop
	Work
	Rest
)

type Pomodoro struct {
	ticker   time.Ticker
	duration time.Duration
	handler  UpdateTickerView
	start    time.Time
	status   PomodoroStatus
}

func (p *Pomodoro) Tick() {
	fmt.Printf("status:%v \n", p.status)
	if p.status == 0 {
		p.status = Work
		p.Run()
	} else if p.status == Work {
		p.status = Rest
		p.Run()
	} else if p.status == Rest {
		p.status = Work
		p.Run()
	}

}

func (p *Pomodoro) Run() {
	fmt.Printf("start0\n")

	p.duration = 0
	p.start = time.Now()

	p.ticker = *time.NewTicker(tickStep)
	var duration time.Duration

	if p.status == Work {
		duration = workDuration
	} else if p.status == Rest {
		duration = restDuration
	}

	fmt.Printf("status:%v, duration:%v \n", p.status, duration)

	go func() {
		for t := range p.ticker.C {
			d := t.Sub(p.start)
			p.handler(d.String())
			if d > duration {
				break
			}
		}
	}()
}

func (p *Pomodoro) Register(view UpdateTickerView) {
	p.handler = view
}

type UpdateTickerView func(value string)
