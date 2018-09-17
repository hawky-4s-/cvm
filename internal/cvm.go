package internal

import (
	"github.com/hawky-4s-/cvm/pkg/internal"
	log "github.com/sirupsen/logrus"
)

func readInitialConfig() (*internal.Config, error) {
	// TODO: merging of configuration or introduce precedence of configurations
	// 1. cmd line
	// 2. env vars
	// 3. local cfg
	// 4. remote cfg
	// read global configuration
	var cfg *internal.Config
	localCfg, err := internal.GetLocalConfiguration("cvm.yaml")
	if err != nil {
		return nil, err
	}
	cfg = localCfg
	remoteCfg, err := internal.GetRemoteConfiguration("", "", "")
	if err != nil {
		return nil, err
	}
	cfg = remoteCfg

	return cfg, nil
}

type Cvm struct {
	config   *internal.Config
	ctx      *internal.CvmContext
	executor *internal.Executor
}

func (c *Cvm) Create(server string, version string) {
	// CREATE TASKS
	// SCHEDULE TASKS TO EXECUTOR
	// EXECUTE TASKS
}

func NewCvm() (*Cvm, error) {
	logger := log.New()

	cfg, err := readInitialConfig()
	if err != nil {
		return nil, err
	}

	cvm := new(Cvm)

	cvm.config = cfg
	cvm.executor = internal.NewExecutor()
	cvm.ctx = internal.NewCvmContext(logger)

	return cvm, nil
}

// init global config
// read user config (env etc, or interactive)
// construct Cvm object with
