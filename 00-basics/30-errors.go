package main

import "errors"
import "fmt"

type WizardLowPowerError struct {
	spellPower    int
	minSpellPower int
}

func (error *WizardLowPowerError) Error() string {
	return fmt.Sprintf(
		"ERROR: Not enough spell power (have %d but %d needed)",
		error.spellPower,
		error.minSpellPower,
	)
}

type Wizard struct {
	spellPower int
}

func (wizard *Wizard) cast(spellCost int) error {
	if spellCost == 666 {
		return errors.New("ERROR: This spell is forbidden")
	}

	if wizard.spellPower-spellCost < 0 {
		return &WizardLowPowerError{
			wizard.spellPower,
			spellCost,
		}
	}

	wizard.spellPower -= spellCost
	fmt.Println("Casting spell with power:", spellCost, ", power left:", wizard.spellPower)

	return nil

}

func main() {
	superWizard := &Wizard{10000}
	err := superWizard.cast(666)
	if err != nil {
		fmt.Println(err.Error())
	}

	wizard := &Wizard{10}
	for {
		err := wizard.cast(3)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
