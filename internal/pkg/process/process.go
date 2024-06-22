package process

import "fmt"

var _ fmt.Stringer = (ProcessName)("")
var _ Nameable = (ProcessNameFunc)(func() ProcessName { panic("not implemented") })

type ProcessName string
type ProcessNameFunc func() ProcessName

// Processor common interface
// to run and stop a process
type Processor interface {
	Run()
	Stop()
}

// Nameable returns the process name
type Nameable interface {
	Name() ProcessName
}

// NameableProcessor like  Processor just
// has the process name
type NameableProcessor interface {
	Nameable
	Processor
}

func (p ProcessName) String() string {
	return string(p)
}

func (pnf ProcessNameFunc) Name() ProcessName {
	return pnf()
}

// NewNameable returns Nameable interface
// ProcessNameFunc is implementing it
// NOTE: I try to avoid casting because it's decreasing performance
func NewNameable(processName string) ProcessNameFunc {
	pName := ProcessName(processName)
	return func() ProcessName {
		return pName
	}
}

// newNameableWithCastForTest 10 times slower than NewNameable
// the cast to ProcessNameFunc has an impact
// for test only
func newNameableWithCastForTest(processName string) Nameable {
	pName := ProcessName(processName)
	return ProcessNameFunc(func() ProcessName {
		return pName
	})
}

// newNameableWithCastForTest 10 times slower than NewNameable
// the cast to ProcessNameFunc has an impact
// for test only
func newNameableWithCastForTest2(processName string) Nameable {
	pName := ProcessName(processName)
	var res ProcessNameFunc = func() ProcessName {
		return pName
	}
	return res
}
