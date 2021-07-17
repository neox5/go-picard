package scheduler

type Scheduler []func() error

func (s Scheduler) Run() error {
	for _, f := range s {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
