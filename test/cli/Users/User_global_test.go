package cli_user_test

import (
	_ "github.com/KewaiiGamer/proxmox-api-go/cli/command/commands"
	cliTest "github.com/KewaiiGamer/proxmox-api-go/test/cli"
	"testing"
)

func Test_User_List(t *testing.T) {
	Test := cliTest.Test{
		Expected: `"userid":"root@pam"`,
		ReqErr:   false,
		Contains: true,
		Args:     []string{"-i", "list", "users"},
	}
	Test.StandardTest(t)
}
