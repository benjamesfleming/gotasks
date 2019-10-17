package actions_test

import (
	"testing"

	"github.com/benjamesfleming/gotasks/actions"
	"github.com/gobuffalo/suite"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{
		Action: suite.NewAction(actions.App()),
	}
	suite.Run(t, as)
}
