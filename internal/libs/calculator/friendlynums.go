package calculator

import (
	"errors"
	"log"
	"math"
	"strconv"
)

func sumOfDividers(number int) int {
	sum := 0

	for i := 1; i < number/2+1; i++ {
		if number%i == 0 {
			sum += i
		}
	}

	return sum
}

func GetAmicableNumber(input string) (string, bool, error) {
	inputAsNumber, err := strconv.Atoi(input)
	switch {
	case err != nil:
		log.Fatal(err)
	case inputAsNumber <= 0:
		log.Fatal(errors.New("number should be greater than 0"))
	case inputAsNumber > math.MaxInt32:
		log.Fatal(errors.New("number should be lesser than " + strconv.Itoa(math.MaxInt32)))
	}

	possibleFriendlyNumber := sumOfDividers(int(inputAsNumber))

	return strconv.FormatInt(int64(possibleFriendlyNumber), 10), sumOfDividers(possibleFriendlyNumber) == inputAsNumber, nil
}
