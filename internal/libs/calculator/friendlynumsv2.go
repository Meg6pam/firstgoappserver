package calculator

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

func sumOfDividersv2(number int64) int64 {
	var sum int64 = 0
	var lastCandidate int64 = 0
	var candidate int64
	for candidate = 2; candidate <= int64(math.Sqrt(float64(number))); candidate++ {
		if number%candidate == 0 {
			sum += candidate + number/candidate
			lastCandidate = candidate
		}
	}

	if lastCandidate != 0 && number/lastCandidate == lastCandidate {
		sum -= lastCandidate
	}

	return sum
}

func GetAmicableNumberv2(input string) (string, bool, error) {
	t0 := time.Now()
	t1 := time.Now()
	inputAsNumber, err := strconv.ParseInt(input, 10, 0)
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

	var possibleFriendlyNumber int64 = sumOfDividersv2(int64(inputAsNumber))

	return strconv.FormatInt(int64(possibleFriendlyNumber), 10), sumOfDividersv2(possibleFriendlyNumber) == inputAsNumber, nil
}
