package strand

func ToRNA(dna string) string {
	if len(dna) <= 0{
        return ""
    }
    rna := ""
    for i := 0; i < len(dna); i++ {
        switch dna[i]{
            case 'G':
                 rna += "C"
            case 'C':
                 rna += "G"
            case 'T':
                 rna += "A"
            case 'A':
                 rna += "U"
        }
    }
    return rna
}
