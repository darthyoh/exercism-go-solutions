package protein

import "errors"

//ErrStop : get a STOP signal Codon
var ErrStop = errors.New("Stop")

//ErrInvalidBase : get an unknow Codon
var ErrInvalidBase = errors.New("Invalid")

//FromCodon gives a protein from a Codon
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

//FromRNA gives proteins array from RNA sequence
func FromRNA(rna string) ([]string, error) {
	results := make([]string, 0)
	for i := 0; i <= len(rna)-3; i += 3 {
		if codon, err := FromCodon(rna[i : i+3]); err == nil {
			results = append(results, codon)
		} else {
			if err == ErrStop {
				return results, nil
			}
			return results, ErrInvalidBase
		}
	}
	return results, nil
}
