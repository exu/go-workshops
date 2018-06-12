package main

import "errors"
import "fmt"

// WizardLowPowerError - our custom error like exception in other OO langs
type WizardLowPowerError struct {
	spellPower    int
	minSpellPower int
}

// need to implement `Error() string` method
func (error *WizardLowPowerError) Error() string {
	return fmt.Sprintf(
		"ERROR: Not enough mana (have %d but %d needed)",
		error.spellPower,
		error.minSpellPower,
	)
}

// Wizard - Our magic struct
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
	fmt.Println("Casting spell with power:", spellCost, ", mana left:", wizard.spellPower)

	return nil

}

func main() {
	// first instance will return error when we will cast spell with power value of 666
	superWizard := Wizard{10000}
	err := superWizard.cast(666)
	if err != nil {
		fmt.Println(err.Error())
	}

	// now we will be casting spells as long as error will occur.
	wizard := Wizard{100}
	for {
		err := wizard.cast(3)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
