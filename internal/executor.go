package internal

import (
	"container/list"
)

type ExecCommand func(ctx *CvmContext) (RollbackCommand, error)
type RollbackCommand func(crx *CvmContext) error

type CommandType int

const (
	ShellCommand CommandType = 1 << iota
	DownloadCommand
	ProvisionCommand
	DockerCommand
)

type Executor struct {
	//queue []ExecCommand
	queue         *list.List
	rollbackQueue *list.List
	doRollback    bool
}

func (e *Executor) Empty() bool {
	return e.queue.Len() == 0
}

func (e *Executor) schedule(cmd ExecCommand) {
	e.queue.PushBack(cmd)
}

func (e *Executor) execute(ctx *CvmContext) *ExecutionError {
	var executionErr error
	var rollbackErr error

	defer func() {
		e.queue.Init()
		e.rollbackQueue.Init()
	}()

	for element := e.queue.Front(); element.Next() != nil; element = element.Next() {
		cmd := element.Value.(ExecCommand)

		ctx.Log.Debug("executing cmd")
		rollbackCmd, err := cmd(ctx)
		e.rollbackQueue.PushFront(rollbackCmd)
		if err != nil {
			ctx.Log.Error("cmd execution failed")
			executionErr = err
			break
		}
	}
	if executionErr == nil {
		return nil
	}

	if executionErr != nil && e.doRollback {
		ctx.Log.Info("starting rollback of cmd execution")
		for element := e.rollbackQueue.Front(); element.Next() != nil; element = element.Next() {
			cmd := element.Value.(RollbackCommand)

			ctx.Log.Info("executing rollback cmd")
			err := cmd(ctx)
			if err != nil {
				ctx.Log.Error("rollback cmd execution failed")
				rollbackErr = err
				break
			}
		}
	}

	return NewExecutionError(executionErr, rollbackErr)
}

func NewExecutor() *Executor {

	return &Executor{
		queue:         list.New(),
		rollbackQueue: list.New(),
		doRollback:    true,
	}
}

type CommandFactory struct {
}

func (f CommandFactory) createDownloadCommand(filepath, url, username, password string) ExecCommand {
	return nil
}

func (f CommandFactory) createProvisionCommand(cmd ...string) ExecCommand {
	return nil
}

func (f CommandFactory) createShellCommand(cmd string) ExecCommand {
	return nil
}
