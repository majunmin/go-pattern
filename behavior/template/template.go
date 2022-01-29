// Created By: junmin.ma
// Description: <description>
// Date: 2022-01-29 18:38
package template

// WorkerInterface define an interface of worker.
type WorkerInterface interface {
	GetUp()
	EatBreakfast()
	GoToWork()
	Work()
	ReturnHome()
	Relax()
	Sleep()
}

// defines a worker
type Worker struct {
	WorkerInterface
}

func NewWorker(workerInterface WorkerInterface) *Worker {
	return &Worker{WorkerInterface: workerInterface}
}

// DailyRoutine template method
func (w *Worker) DailyRoutine() {
	w.GetUp()
	w.EatBreakfast()
	w.GoToWork()
	w.Work()
	w.ReturnHome()
	w.Relax()
	w.Sleep()
}

// Postman  concrete Worker
type Postman struct {
}

func (p *Postman) GetUp() {
	panic("implement me")
}

func (p *Postman) EatBreakfast() {
	panic("implement me")
}

func (p *Postman) GoToWork() {
	panic("implement me")
}

func (p *Postman) Work() {
	panic("implement me")
}

func (p *Postman) ReturnHome() {
	panic("implement me")
}

func (p *Postman) Relax() {
	panic("implement me")
}

func (p *Postman) Sleep() {
	panic("implement me")
}
