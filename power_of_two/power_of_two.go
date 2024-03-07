package power_of_two

/*
Task: return bool of the int is a power of 2
Use bitwise operations exclusively

# Examples

0     -> false
1     -> true
10    -> true
11    -> false
100   -> true
1000  -> true
10000 -> true

Conclusion: return true if only 1 bit in the int is on

# Algorithm
if i & 1 == 1 -> last bit in i is on

*/

func IsPowerOfTwo(i int) bool {
	previousBitOn := false
	for i > 0 {
		if i&1 == 1 {
			if previousBitOn {
				return false
			} else {
				previousBitOn = true
			}
		}
		i >>= 1
	}
	return previousBitOn // Benchmark with if logic at start
}
