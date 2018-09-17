package internal

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func createTestCvmContext() *CvmContext {
	return NewCvmContext(logrus.New())
}

func TestNewExecutor(t *testing.T) {
	cvmContext := createTestCvmContext()

	executor := NewExecutor()

	fooCmd := func(ctx *CvmContext) (RollbackCommand, error) {
		s := "foo"
		fmt.Println(s)
		return func(ctx *CvmContext) error {
			split := strings.Split(s, "")
			sort.Strings(split)
			fmt.Println(s)
			fmt.Println(split)
			return nil
		}, nil
	}

	barCmd := func(ctx *CvmContext) (RollbackCommand, error) {
		s := "bar"
		fmt.Println(s)
		return func(ctx *CvmContext) error {
			split := strings.Split(s, "")
			sort.Strings(split)
			fmt.Println(s)
			fmt.Println(split)
			return nil
		}, nil
	}

	errCmd := func(ctx *CvmContext) (RollbackCommand, error) {
		fmt.Println("errors out")
		err := errors.New("just dying")
		return func(crx *CvmContext) error {
			if err != nil {
				fmt.Println("I am alive")
			}
			return nil
		}, err
	}

	executor.schedule(fooCmd)
	executor.schedule(errCmd)
	executor.schedule(barCmd)
	execErrors := executor.execute(cvmContext)
	assert.NotNil(t, execErrors)
	assert.Len(t, execErrors.executionErrors, 1)

	assert.True(t, executor.doRollback, "doRollback should be true")
	assert.Equal(t, 2, executor.rollbackQueue.Len(), "rollback queue should contain all rollback cmds")
	assert.Equal(t, 3, executor.queue.Len(), "queue should still contain the cmds")
}
