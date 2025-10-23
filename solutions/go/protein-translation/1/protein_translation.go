package protein
import (
	"errors"
)
var (
	ErrStop        = errors.New("stop codon encountered")
	ErrInvalidBase = errors.New("invalid codon")
)
// codonTable maps RNA codons to their corresponding amino acids
var codonTable = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", 
    "UUC": "Phenylalanine",
	"UUA": "Leucine", 
    "UUG": "Leucine",
	"UCU": "Serine", 
    "UCC": "Serine", 
    "UCA": "Serine", 
    "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	var proteins []string
	
	// Process the RNA string in chunks of 3 characters (codons)
	for i := 0; i < len(rna); i += 3 {
		// Ensure we don't go beyond the string length
		if i+3 > len(rna) {
			return nil, ErrInvalidBase
		}
		
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if err != nil {
			if err == ErrStop {
				// For STOP codons, we just break out of the loop
				break
			}
			return nil, err
		}
		
		// Stop translation if we encounter a STOP codon
		if protein == "STOP" {
			break
		}
		
		proteins = append(proteins, protein)
	}
	
	return proteins, nil
    //panic("Please implement the FromRNA function")
}

func FromCodon(codon string) (string, error) {
	protein, exists := codonTable[codon]
	if !exists {
		return "", ErrInvalidBase
	}
	
	if protein == "STOP" {
		return "", ErrStop
	}
	
	return protein, nil
    //panic("Please implement the FromCodon function")
}
