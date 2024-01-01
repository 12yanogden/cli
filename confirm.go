package cli

import "fmt"

type ValidResponses struct {
	Acceptables []string
	Rejectables []string
}

func Confirm(msg string) bool {
	var response string
	validResponses := ValidResponses{
		Acceptables: []string{"Y", "y"},
		Rejectables: []string{"N", "n"},
	}
	var confirmation bool

	fmt.Print(msg)

	fmt.Scanln(&response)

	for !validResponses.Has(response) {
		fmt.Printf("unrecognized response: '%s'. Try again:\n", response)
		fmt.Scanln(&response)
	}

	if validResponses.Accepts(response) {
		confirmation = true
	} else if validResponses.Rejects(response) {
		confirmation = false
	}

	return confirmation
}

func (validResponses ValidResponses) Has(response string) bool {
	return validResponses.Accepts(response) || validResponses.Rejects(response)
}

func (validResponses ValidResponses) Accepts(response string) bool {
	for _, acceptable := range validResponses.Acceptables {
		if response == acceptable {
			return true
		}
	}

	return false
}

func (validResponses ValidResponses) Rejects(response string) bool {
	for _, rejectable := range validResponses.Rejectables {
		if response == rejectable {
			return true
		}
	}

	return false
}
