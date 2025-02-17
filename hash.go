package task

import (
	"fmt"

	"github.com/sciabarracom/task/v3/internal/hash"
	"github.com/sciabarracom/task/v3/internal/slicesext"
	"github.com/sciabarracom/task/v3/taskfile/ast"
)

func (e *Executor) GetHash(t *ast.Task) (string, error) {
	r := slicesext.FirstNonZero(t.Run, e.Taskfile.Run)
	var h hash.HashFunc
	switch r {
	case "always":
		h = hash.Empty
	case "once":
		h = hash.Name
	case "when_changed":
		h = hash.Hash
	default:
		return "", fmt.Errorf(`ops: invalid run "%s"`, r)
	}
	return h(t)
}
