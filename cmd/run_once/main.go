package main

import (
	"fmt"
	"github.com/juju/fslock"
	"log"
	"os"
	"os/exec"
)

func main() {
	err := Main()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Main() error {
	args := os.Args
	if len(args) < 2 {
		return fmt.Errorf("usage: run_once cmd args")
	}
	cmd := args[1]
	args = args[2:]

	lockName, err := getLockPath(cmd, args)
	if err != nil {
		return err
	}

	lock := fslock.New(lockName)
	err = lock.TryLock()
	if err != nil {
		return err
	}
	defer func() {
		_ = lock.Unlock()
	}()

	command := exec.Command(cmd, args...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	err = command.Run()
	if err != nil {
		return err
	}

	return nil
}
