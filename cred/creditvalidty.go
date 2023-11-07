package creditvalidty

import (
	"errors"
)

func CheckCard(cardNum *[]int) error {
	verifyDoubleEvenDigi(cardNum)
	verifySumDoubleEvenDigi(cardNum)
	check := verifyMod10Total(cardNum)

	if !check {
		err := errors.New("Card Invalid")
		return err
	}
	return nil

}

func verifyDoubleEvenDigi(cardNum *[]int) {
	for i := 0; i < 11; i++ {
		if i%2 != 0 {
			(*cardNum)[i] *= 2
		}

	}

}

func verifySumDoubleEvenDigi(cardNum *[]int) {
	for i := 0; i < 11; i++ {
		if (*cardNum)[i] > 9 {
			(*cardNum)[i] = (*cardNum)[i]/10 + (*cardNum)[i]%10
		}

	}

}

func verifyMod10Total(cardNum *[]int) bool {
	sum := 0
	for i := 0; i < 11; i++ {
		sum += (*cardNum)[i]
	}

	if sum%10 != 0 {
		return false
	}
	return true
}
