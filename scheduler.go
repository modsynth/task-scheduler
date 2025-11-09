package scheduler

import (
	"github.com/robfig/cron/v3"
	"log"
)

type Scheduler struct {
	cron *cron.Cron
}

func New() *Scheduler {
	return &Scheduler{
		cron: cron.New(),
	}
}

func (s *Scheduler) AddJob(spec string, cmd func()) error {
	_, err := s.cron.AddFunc(spec, cmd)
	if err != nil {
		log.Printf("Failed to add job: %v", err)
		return err
	}
	return nil
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}
