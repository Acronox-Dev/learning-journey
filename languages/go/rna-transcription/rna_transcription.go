package rnatranscription

func getNucleotideComplement(n rune) rune {
	switch n {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	default:
		return 0
	}
}

func ToRNA(dna string) string {
	rna := make([]rune, len(dna))

	for i := 0; i < len(dna); i++ {
		rna[i] = getNucleotideComplement(rune(dna[i]))
	}

	return string(rna)
}
