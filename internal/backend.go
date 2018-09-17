package internal

// Backend is the interface to interact with the different provisioning backends supported by CVM.
type Backend interface {
}

type LocalBackend struct {
}

type DockerBackend struct {
}
