package main

import "gitTest/sportsman"

func main() {
	w, h, a, err := sportsman.GetCharacter()
	if err != nil {
		return
	}
	dayLimit := sportsman.CountRequiredCalories(w, h, a)
	sportsman.CheckBmr(dayLimit, 5000)
}
