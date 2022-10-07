package main

import "fmt"

//lista de tareas
type taskList struct {
	tasks []*task
}

func (tl *taskList) appendTask(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

type task struct {
	name  string
	state bool
	desc  string
}

func (t *task) showTask() {
	fmt.Println(t.name, " ", t.desc, " ", t.state)
}

func (t *task) completeTask() {
	t.state = true
}

func main() {
	task1 := task{name: "Estudiar Go", state: false, desc: "Estudiar dos horas de Go"}
	task2 := task{name: "Estudiar Node", state: true, desc: "Estudiar una hora de Node"}
	task3 := task{name: "Estudiar Vue", state: true, desc: "Estudiar dos hora de Vue"}
	//task1.showTask()
	//task1.completeTask()
	//task1.showTask()

	lista := taskList{}
	lista.appendTask(&task1) // envio task1 como referencia para no clonar
	lista.appendTask(&task2) // envio task1 como referencia para no clonar
	lista.appendTask(&task3) // envio task1 como referencia para no clonar
	fmt.Println(lista)       //mire que viene dos llaves es por eso que utiliza elk punto en el for

	for i, task := range lista.tasks {
		fmt.Println(i, task.name)
	}
}
