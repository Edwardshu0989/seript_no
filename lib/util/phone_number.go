package util

import (
	"fmt"
	"loanmarket-server/lib/i18n"
	"strings"
)

type SplitPhone struct {
	AreaCode string
	Phone    string
}

func ParsePhoneNumber(phone string) (SplitPhone, error) {
	arr := strings.Split(phone, "-")
	if len(arr) == 2 {
		return SplitPhone{
			AreaCode: arr[0],
			Phone:    arr[1],
		}, nil
	}
	return SplitPhone{}, fmt.Errorf(i18n.PhoneFormatError)
}
