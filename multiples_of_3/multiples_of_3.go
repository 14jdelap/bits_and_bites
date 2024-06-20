package multiples_of_three

/*
Examples

3:  11
6:  110
9:  1001
12: 1100
15: 1111
18: 10010
21: 10101
24: 11000
27: 11011
30: 11110
33: 100001

Modulo mathematics: r = x - (y * (x/y))
										r = 33 - (3 * (33/3)) = 0

3: 11


*/

func IsMultipleOfThree(i int) bool {
	return 9%3 == 0
}
