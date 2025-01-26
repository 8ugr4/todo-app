package main

import (
	"fmt"
	"strings"
)

type Task struct {
	name   string
	status bool
}

type Tasks map[int]Task

type FinishedTasks map[int]Task

//type taskMaps interface {
//	Tasks
//	FinishedTasks
//}

func (ft *FinishedTasks) removeFinishedTasks(ts *Tasks) {
	for id, task := range *ts {
		if task.status == true {
			ft.addFinishedTasks(task)
			delete(*ts, id)
		}
	}
}

func (ft *FinishedTasks) addFinishedTasks(t Task) {
	numberOfFinishedTasks := len(*ft)
	(*ft)[numberOfFinishedTasks] = t
}

func (ft *FinishedTasks) clear() {
	for id := range *ft {
		delete(*ft, id)
	}
}

func (ft *FinishedTasks) showFinishedTasks() {
	nameWidth := findLongestName(*ft)
	totalWidth := nameWidth + maxIDWidth + maxStatusWidth + 4
	halfSpace := strings.Repeat("-", totalWidth/2-len(fmt.Sprintf("finished-tasks"))/2+2)
	hyphens := calculateHyphenFormat(nameWidth)

	fmt.Printf("\n+%s%s%s+\n", halfSpace, "-finished-tasks", halfSpace)
	fmt.Printf("| %-*s | %-*s | %-*s |\n", maxIDWidth, "ID",
		nameWidth, "NAME", maxStatusWidth, "STATUS")
	fmt.Printf("| %-*s | %-*s | %-*v |\n", maxIDWidth, hyphens[ID], nameWidth, hyphens[name], maxStatusWidth, hyphens[status])

	for id, task := range *ft {
		fmt.Printf("| %-*d | %-*s | %-*v |\n", maxIDWidth, id, nameWidth, task.name, maxStatusWidth, "done")
	}
	fmt.Printf("+%s+\n", strings.Repeat("-", totalWidth+4))
}

func (t *Task) createTask(taskName string, taskStatus bool) {
	t.status = taskStatus
	t.name = taskName
}

func (ts *Tasks) addTask(t Task) {
	numberOfTasks := len(*ts)
	(*ts)[numberOfTasks] = t
}

func (ts *Tasks) deleteTask(taskID int) {
	for id := range *ts {
		if id == taskID {
			delete(*ts, id)
		}
	}
}

func findLongestName(value interface{}) int {
	var nameWidth int

	switch ts := value.(type) {
	case Tasks:
		nameWidth = len(ts[0].name)
		for _, t := range ts {
			if nameWidth < len(t.name) {
				nameWidth = len(t.name)
			}
		}
	case FinishedTasks:
		nameWidth = len(ts[0].name)
		for _, t := range ts {
			if nameWidth < len(t.name) {
				nameWidth = len(t.name)
			}
		}
	}
	if nameWidth%2 == 0 {
		nameWidth += 1
	}

	return nameWidth
}

const (
	maxIDWidth     = 2
	maxStatusWidth = 6
	ID             = "id"
	status         = "status"
	name           = "name"
)

var hyphenFormat map[string]string

func calculateHyphenFormat(nameWidth int) map[string]string {
	nameHyphen := strings.Repeat("-", nameWidth)
	statusHyphen := strings.Repeat("-", maxStatusWidth)
	idHyphen := strings.Repeat("-", maxIDWidth)

	hyphenFormat = make(map[string]string)

	hyphenFormat["name"] = nameHyphen
	hyphenFormat["status"] = statusHyphen
	hyphenFormat["id"] = idHyphen

	return hyphenFormat
}

func (ts *Tasks) showTasks() {
	nameWidth := findLongestName(*ts)
	totalWidth := nameWidth + maxIDWidth + maxStatusWidth + 4

	halfSpace := strings.Repeat("-", totalWidth/2)
	hyphens := calculateHyphenFormat(nameWidth)

	fmt.Printf("\n+%s%s%s+\n", halfSpace, "tasks", halfSpace)
	fmt.Printf("| %-*s | %-*s | %-*s |\n", maxIDWidth, "ID",
		nameWidth, "NAME", maxStatusWidth, "STATUS")
	fmt.Printf("| %-*s | %-*s | %-*v |\n", maxIDWidth, hyphens[ID], nameWidth, hyphens[name], maxStatusWidth, hyphens[status])

	for id, task := range *ts {
		fmt.Printf("| %-*d | %-*s | %-*v |\n", maxIDWidth, id, nameWidth, task.name, maxStatusWidth, task.status)

	}
	fmt.Printf("+%s+\n", strings.Repeat("-", totalWidth+4))
}

func (ts *Tasks) updateTaskName(taskID int, taskName string) {
	for id, task := range *ts {
		if id == taskID {
			task.name = taskName
		}
	}
}

func (ts *Tasks) updateTaskStatus(taskID int, taskStatus bool) {
	for id, task := range *ts {
		if id == taskID {
			task.status = taskStatus
		}
	}
}

func main() {
	ts := Tasks{}
	t := Task{}
	ft := FinishedTasks{}

	t.createTask("create struct", true)
	ts.addTask(t)

	t.createTask("study ech", false)
	ts.addTask(t)

	t.createTask("write CRUD operations in golang", true)
	ts.addTask(t)

	t.createTask("asdsfadgdfgggggggggggggggg", false)
	ts.addTask(t)

	ts.showTasks()
	ft.removeFinishedTasks(&ts)
	ts.showTasks()
	ft.showFinishedTasks()
}
