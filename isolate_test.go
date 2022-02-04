package isolate

import (
	"fmt"
	"os"
	"testing"
)

func TestInitAndCleanup(t *testing.T) {
	sandbox, err := New(0, true)
	if err != nil {
		t.Error(err)
		return
	}

	if sandbox.Id != "0" {
		t.Error("Wrong ID")
		return
	}

	err = sandbox.CleanUp()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("success")
}

func TestRun(t *testing.T) {
	sandbox, err := New(1, true)
	if err != nil {
		t.Error(err)
		return
	}
	defer sandbox.CleanUp()

	os.WriteFile(sandbox.Path+"/box/test.py", []byte("print(sum(map(int, input().split())))"), 0664)
	cmd := sandbox.Run("/usr/bin/python3", []string{"test.py"},
		TimeLimit(1), MemLimit(1024*16),
		Stdout("stdout.txt"),
	)

	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Fprintln(stdinPipe, 1000, 24)
	stdinPipe.Close()

	err = cmd.Start()
	if err != nil {
		t.Error(err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		t.Error(err)
		return
	}

	stdout, err := os.ReadFile(sandbox.Path + "/box/stdout.txt")
	if err != nil {
		t.Error(err)
		return
	}
	if string(stdout) != "1024\n" {
		fmt.Println(string(stdout))
		t.Error("output incorrect")
		return
	}

	t.Log("success")
}
