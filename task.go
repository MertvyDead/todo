package main

type Task struct {
	Title  string
	IsDone bool
}

func (t *Task) done() {
	t.IsDone = true
}
