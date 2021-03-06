package action

import (
	"github.com/dodo-cli/dodfile-syntax/pkg/state"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/dockerfile/dockerfile2llb"
)

type ScriptAction struct {
	Script string
	User   string
	Cwd    string
}

func (a *ScriptAction) Execute(base llb.State) llb.State {
	s := state.FromLLB(defaultBaseImage, base)

	if len(a.User) > 0 {
		s.User(a.User)
	}

	if len(a.Cwd) > 0 {
		s.Cwd(a.Cwd)
	}

	s.Sh(a.Script)

	return s.Get()
}

func (*ScriptAction) UpdateMetadata(_ *dockerfile2llb.Image) {
}
