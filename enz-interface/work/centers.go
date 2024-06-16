package work

type EnzoWorkcenter interface {
	Observable
	Id() string
	Init()
	Queue(EnzoWorkItem)
}
