package main

type Task struct {
	ID     int
	Title  string
	IsDone bool
}

func (t *Task) done() {
	t.IsDone = true
}
func (t *Task) newID(id int) {
	t.ID = id
}
