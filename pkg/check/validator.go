package check

import (
	"errors"
	"regexp"
	"time"
)

func ValidateAge(age int, birthday string) error {
	birthDate, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return errors.New("invalid birthday format")
	}

	currentTime := time.Now()
	calculatedAge := currentTime.Year() - birthDate.Year()
	if currentTime.YearDay() < birthDate.YearDay() {
		calculatedAge--
	}
	
	if age != calculatedAge {
		return errors.New("age does not match the birthday")
	}

	return nil
}

func ValidatePhone(phone string) bool {
	return regexp.MustCompile(`^\+998[0-9]{9}$`).MatchString(phone)
}

func ValidateGmail(gmail string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@gmail.com$`).MatchString(gmail)
}

func ValidatePassword(password string) bool {
	return regexp.MustCompile(`^[A-Z0-9!@#$+]{9}`).MatchString(password)
}