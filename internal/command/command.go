package command

import (
	"os"
	"strings"

	"slices"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CommandExecutor struct {
	DB *gorm.DB
}

func NewCommandExecutor(db *gorm.DB) *CommandExecutor {
	return &CommandExecutor{DB: db}
}

func (ce *CommandExecutor) Execute(logger *logrus.Logger) bool {
	args := os.Args[1:]
	run := true
	hasFlag := false

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			hasFlag = true
			switch arg {
			case "--run":
				run = true
			}
		}
	}

	if hasFlag && !contains(args, "--run") {
		run = false
	}

	return run
}

func contains(slice []string, target string) bool {
	return slices.Contains(slice, target)
}