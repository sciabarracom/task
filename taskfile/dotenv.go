package taskfile

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/sciabarracom/task/v3/internal/compiler"
	"github.com/sciabarracom/task/v3/internal/filepathext"
	"github.com/sciabarracom/task/v3/internal/templater"
	"github.com/sciabarracom/task/v3/taskfile/ast"
)

func Dotenv(c *compiler.Compiler, tf *ast.Taskfile, dir string) (*ast.Vars, error) {
	if len(tf.Dotenv) == 0 {
		return nil, nil
	}

	vars, err := c.GetTaskfileVariables()
	if err != nil {
		return nil, err
	}

	env := &ast.Vars{}
	cache := &templater.Cache{Vars: vars}

	for _, dotEnvPath := range tf.Dotenv {
		dotEnvPath = templater.Replace(dotEnvPath, cache)
		if dotEnvPath == "" {
			continue
		}
		dotEnvPath = filepathext.SmartJoin(dir, dotEnvPath)

		if _, err := os.Stat(dotEnvPath); os.IsNotExist(err) {
			continue
		}

		envs, err := godotenv.Read(dotEnvPath)
		if err != nil {
			return nil, err
		}
		for key, value := range envs {
			if ok := env.Exists(key); !ok {
				env.Set(key, ast.Var{Value: value})
			}
		}
	}

	return env, nil
}
