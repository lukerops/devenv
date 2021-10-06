package model

type Container struct {
	ID       string
	Names    []string
	Labels   map[string]string
	State    string
	Exited   bool
	ExitCode int8
	Image    string
	ImageID  string
	Mounts   []string
	Networks []string
	Ports    []string
}
