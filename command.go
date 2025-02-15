package main

import "fmt"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandsMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandsMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if cmdFunc, ok := c.commandsMap[cmd.name]; ok {
		cmdFunc(s, cmd)
	} else {
		return fmt.Errorf("%s - command not registered", cmd.name)
	}
	return nil
}
