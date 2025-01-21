package internal

type Task struct {
	ID     int
	Title  string
	IsDone bool
}

func (t *Task) done() {
	t.IsDone = true
}
