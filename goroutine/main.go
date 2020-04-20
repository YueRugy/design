package main

import (
	"fmt"
	"time"
)

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{
		f: f,
	}
}

func (t *Task) Execute() {
	err := t.f()
	if err != nil {
		fmt.Println(err)
	}
}

type Pool struct {
	EntryChan chan *Task
	JobChan   chan *Task
	max       int
}

func NewPool(max int) *Pool {
	return &Pool{
		EntryChan: make(chan *Task, 1000),
		JobChan:   make(chan *Task, 1000),
		max:       max,
	}
}

func (p *Pool) worker(workId int) {
	for t := range p.JobChan {
		t.Execute()
		fmt.Println("work_id", workId, "执行完毕")
	}
}

func (p *Pool) run() {
	for i := 0; i < p.max; i++ {
		go p.worker(i)
	}

	for task := range p.EntryChan {
		p.JobChan <- task
	}
}

func main() {

	task := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	p := NewPool(4)
	go func() {
		for {
			p.EntryChan <- task
		}
	}()
	p.run()

}
