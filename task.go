package todo

type Task struct {
	id     int
	title  string
	isDone bool
}

func (t *Task) done() {
	t.isDone = true
}
