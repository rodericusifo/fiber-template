package constant

type Environment string

var (
	DEV    = Environment("dev")
	DOCKER = Environment("docker")
)
