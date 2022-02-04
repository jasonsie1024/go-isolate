package isolate

import (
	"os/exec"
	"strconv"
	"strings"
)

type Sandbox struct {
	Id           string
	Path         string
	ControlGroup bool
	Options      Options
}

func mustHaveIsolate() {
	err := exec.Command("isolate", "--version").Run()
	if err != nil {
		panic(err)
	}
}

func (s *Sandbox) CleanUp() error {
	args := []string{"--box-id", s.Id}
	if s.ControlGroup {
		args = append(args, "--cg")
	}

	args = append(args, "--cleanup")

	return exec.Command("isolate", args...).Run()
}

func (s *Sandbox) Run(program string, arguments []string, options ...WithOption) *exec.Cmd {
	for _, wo := range options {
		wo(s)
	}

	args := []string{"--box-id", s.Id}
	if s.ControlGroup {
		args = append(args, "--cg")
	}
	args = append(args, s.Options.BuildArguments()...)
	args = append(args, "--run")
	args = append(args, "--", program)
	args = append(args, arguments...)

	command := exec.Command("isolate", args...)

	return command
}

// Create new sandbox with specified Id and whether to enable Control Group.
func New(Id int, ControlGroup bool) (Sandbox, error) {
	mustHaveIsolate()

	sandbox := Sandbox{
		Id:           strconv.Itoa(Id),
		ControlGroup: ControlGroup,
	}

	args := []string{"--box-id", sandbox.Id}
	if ControlGroup {
		args = append(args, "--cg")
	}
	args = append(args, "--init")

	output, err := exec.Command("isolate", args...).Output()
	if err != nil {
		return Sandbox{}, err
	}

	sandbox.Path = strings.TrimSuffix(string(output), "\n")

	return sandbox, nil
}
