package process_config

type ProcessConfigure interface {
	GetProcessName() string
}

// GetProcessNameFunc implements ProcessConfigure interface
// use function to keep on immutability
type GetProcessNameFunc func() string

func (fn GetProcessNameFunc) GetProcessName() string {
	return fn()
}

func NewProcessConfigure(procName string) ProcessConfigure {

	return GetProcessNameFunc(func() string {
		return procName
	})
}
