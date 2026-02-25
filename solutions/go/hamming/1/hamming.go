package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Both strands have different length")
	}

	diff := 0

	for i := 0; i < len(a); i++ {
		if (a[i] != 'C' && a[i] != 'A' && a[i] != 'G' && a[i] != 'T') ||
			(b[i] != 'C' && b[i] != 'A' && b[i] != 'G' && b[i] != 'T') {
			return 0, errors.New("These letters are invalid")
		} else {
			if a[i] != b[i] {
				diff++
			}
		}
	}

	return diff, nil
}