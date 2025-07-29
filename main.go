package main

import (
	"fmt"
	"math/rand"
	"time"
)

var incorrectAttempts int = 0
var currentOTP string
var lockedUntil time.Time

func getOTP() string {
	otp := fmt.Sprintf("%05d", rand.Intn(100000))
	currentOTP = otp
	return otp
}

func checkOTP(input string) string {
	if input == currentOTP {
		incorrectAttempts = 0
		lockedUntil = time.Time{}
		return "OTP is correct!"
	} else {
		incorrectAttempts++
		if incorrectAttempts < 5 {
			return fmt.Sprintf("Incorrect OTP! You have %d chances left.", 5-incorrectAttempts)
		} else {
			lockedUntil = time.Now().Add(20 * time.Second)
			return "Too many incorrect attempts. Please try again after 20 seconds."
		}
	}
}

func main() {
	fmt.Println("OTP System")

	otp := getOTP()
	fmt.Println("OTP is (for testing purposes):", otp)

	for {
		now := time.Now()

		if now.After(lockedUntil) && !lockedUntil.IsZero() {
			incorrectAttempts = 0
			lockedUntil = time.Time{}
		}

		if now.Before(lockedUntil) {
			time.Sleep(1 * time.Second)
			continue
		}

		var input string
		fmt.Print("Enter OTP: ")
		fmt.Scanln(&input)
		result := checkOTP(input)
		fmt.Println(result)

		if result == "OTP is correct!" {
			break
		}
	}
}