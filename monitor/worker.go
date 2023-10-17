package monitor

type Worker interface {
	Work(input any)
}
