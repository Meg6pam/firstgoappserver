package calculator

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
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
	t0 := time.Now()
	t1 := time.Now()
	inputAsNumber, err := strconv.Atoi(input)
	switch {
	case inputAsNumber == -1:
		{
			t0 = time.Now()
		}
	case inputAsNumber == -3:
		{
			t1 = time.Now()
			fmt.Printf("Elapsed time: %v", t1.Sub(t0))
		}
	case err != nil:
		inputAsNumber = -2
	case inputAsNumber <= 0:
		log.Fatal(errors.New("number should be greater than 0"))
	case inputAsNumber > math.MaxInt64:
		log.Fatal(errors.New("number should be lesser than " + strconv.Itoa(math.MaxInt64)))
	}

	possibleFriendlyNumber := sumOfDividers(int(inputAsNumber))

	return strconv.FormatInt(int64(possibleFriendlyNumber), 10), sumOfDividers(possibleFriendlyNumber) == inputAsNumber, nil
}
