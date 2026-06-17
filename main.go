package main

import (
	"errors"
	"fmt"
	"go-onboarding/calculator"
	"go-onboarding/storage"
)

type UserAccessor interface {
	Save(id int, name string) error
	Get(id int) (storage.User, bool)
}

func main() {
	fmt.Println("catch error")
	store := storage.NewDataStore()
	anlayticEngine := calculator.NewAnalytic()

	err := runBusisnessLogic(store, anlayticEngine)
	if err != nil {
		fmt.Printf("Application Fatal Error: %v\n", err)
	}

}

func runBusisnessLogic(accessor UserAccessor, calc *calculator.Analytic) error {
	err := accessor.Save(10, "Dee")
	if err != nil {
		if errors.Is(err, storage.ErrInvalidName) {
			fmt.Println("⚠️ Validation intercept working perfectly")
			return nil
		}
		return fmt.Errorf("failed to save data due to system error: %w", err)
	}

	name, ok := accessor.Get(10)
	if !ok {
		return fmt.Errorf("user migration missing")
	}
	fmt.Println("Suceesfully retrieved name:", name)

	b, err := calc.CalculateBonus(name.ID)
	if err != nil {

		return fmt.Errorf("failed to process analytics: %w", err)
	}
	fmt.Println("Suceesfully calculated bonus:", b)

	err = accessor.Save(11, "")
	if err != nil {
		if errors.Is(err, storage.ErrInvalidName) {
			fmt.Println("⚠️ Validation intercept working perfectly")
			return nil
		}
		return fmt.Errorf("unexpected system failure: %w", err)
	}

	return nil
}
