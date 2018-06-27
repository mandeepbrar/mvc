package patterns

import (
	"github.com/mandeepbrar/mvc/interfaces"
)

type SimpleCommand struct {
	*Notifier
}

func NewSimpleCommand() *SimpleCommand {
	return &SimpleCommand{NewNotifier()}
}

type MacroCommand struct {
	*Notifier
	subCommands []interfaces.Command
}

func NewMacroCommand() *MacroCommand {
	return &MacroCommand{NewNotifier(), make([]interfaces.Command, 0)}
}

func (mc *MacroCommand) AddSubCommand(comm interfaces.Command) {
	mc.subCommands = append(mc.subCommands, comm)
}

func (mc *MacroCommand) Execute(not interfaces.Notification) {
	for _, comm := range mc.subCommands {
		comm.InitializeNotifier(mc.multitonKey)
		comm.Execute(not)
	}
}
