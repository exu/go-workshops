package main

import "errors"
import "fmt"

// WizardLowPowerError - Nasz customowy błąd
type WizardLowPowerError struct {
	spellPower    int
	minSpellPower int
}

func (error *WizardLowPowerError) Error() string {
	return fmt.Sprintf(
		"ERROR: Not enough mana (have %d but %d needed)",
		error.spellPower,
		error.minSpellPower,
	)
}

// Wizard - Nasz czarodziej
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
	// pierwsza instancja rzuci błędem gdy będziemy
	// rzucać czar o mocy 666
	superWizard := Wizard{10000}
	err := superWizard.cast(666)
	if err != nil {
		fmt.Println(err.Error())
	}

	// i wykorzystanie customowego błędu
	// rzucamy tak długo aż nie wyczerpiemy
	// dostępnej mocy
	wizard := Wizard{100}
	for {
		err := wizard.cast(3)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
