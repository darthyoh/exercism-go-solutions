package perfect

import "errors"

//ErrOnlyPositive throw if number isn't real
var ErrOnlyPositive = errors.New("Only Positive")

//Classification type for number
type Classification int

const (
	//ClassificationDeficient for number < aliquot sum
	ClassificationDeficient Classification = iota
	//ClassificationPerfect for number == aliquot sum
	ClassificationPerfect Classification = iota
	//ClassificationAbundant for number > aliquot sum
	ClassificationAbundant Classification = iota
)

//Classify determines if a number is perfect, abundant, or deficient
func Classify(limit int64) (Classification, error) {
	if limit < 1 {
		return 0, ErrOnlyPositive
	}

	sum := int64(0)

	for i := int64(1); i < limit; i++ {
		if limit%i == 0 {
			sum += i
		}
	}

	switch {
	case sum == limit:
		return ClassificationPerfect, nil
	case sum > limit:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}

}
