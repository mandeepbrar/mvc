package patterns

import (
	"mvc/api"
)

type SimpleCommand struct {
	*Notifier
}

func NewSimpleCommand() *SimpleCommand {
	return &SimpleCommand{NewNotifier()}
}

type MacroCommand struct {
	*Notifier
	subCommands []api.Command
}

func NewMacroCommand() *MacroCommand {
	return &MacroCommand{NewNotifier(), make([]api.Command, 0)}
}

func (mc *MacroCommand) AddSubCommand(comm api.Command) {
	mc.subCommands = append(mc.subCommands, comm)
}

func (mc *MacroCommand) Execute(not api.Notification) {
	for _, comm := range mc.subCommands { 
		comm.InitializeNotifier( mc.multitonKey );
		comm.Execute( not )
	}
}