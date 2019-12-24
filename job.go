package bus

type Job interface {
	execute(params string) (string, error)
}
