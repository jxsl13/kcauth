package auth

import (
	configo "github.com/jxsl13/simple-configo"
	"github.com/manifoldco/promptui"
)

var (
	// DefaultPasswordPrompt is the default configuration for prompting for a password from the user
	DefaultPasswordPrompt = promptui.Prompt{
		Label:       "Password",
		Mask:        '*',
		HideEntered: true,
	}

	// DefaultUsernamePrompt is the default configuration for prompting for a username from the user
	DefaultUsernamePrompt = promptui.Prompt{
		Label:       "Username",
		HideEntered: true,
	}
)

func PromptPassword(outPassword *string) configo.ActionFunc {
	return func() error {
		password, err := DefaultPasswordPrompt.Run()
		if err != nil {
			return err
		}
		*outPassword = password
		return nil
	}
}

func PromptText(outUsername *string) configo.ActionFunc {
	return func() error {
		username, err := DefaultUsernamePrompt.Run()
		if err != nil {
			return err
		}
		*outUsername = username
		return nil
	}
}
