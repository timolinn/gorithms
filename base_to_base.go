package gorithms

// BaseToBase converts value in base base to newBase
func BaseToBase(value string, base, newBase int) string {
	return DecToBase(BaseToDec(value, base), newBase)
}
