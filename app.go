package main

// func main() { //Отзеркаливания массива
//
//		A := []int{1, 2, 3, 4, 5, 6}
//		var n = 0
//		var k = len(A) - len(A)/2
//		for i := len(A) - 1; i >= k; i-- {
//			j := A[n]
//			A[n] = A[i]
//			A[i] = j
//			n++
//		}
//		fmt.Println(A)
//	}
//
// func main() { //Задача на песочные часы
//
//	arr := [6][6]int32{
//		{-9, -9, -9, 1, 1, 1},
//		{0, -9, 0, 4, 3, 2},
//		{-9, -9, -9, 1, 2, 3},
//		{0, 0, 8, 6, 6, 0},
//		{0, 0, 0, -2, 0, 0},
//		{0, 0, 1, 2, 4, 0},
//	}
//	arr2 := [4][4]int32{}
//
//	arr2[0][0] = arr[0][0] + arr[0][1] + arr[0][2] + arr[1][1] + arr[2][0] + arr[2][1] + arr[2][2]
//	arr2[0][1] = arr[0][1] + arr[0][2] + arr[0][3] + arr[1][2] + arr[2][1] + arr[2][2] + arr[2][3]
//	arr2[0][2] = arr[0][2] + arr[0][3] + arr[0][4] + arr[1][3] + arr[2][2] + arr[2][3] + arr[2][4]
//	arr2[0][3] = arr[0][3] + arr[0][4] + arr[0][5] + arr[1][4] + arr[2][3] + arr[2][4] + arr[2][5]
//
//	arr2[1][0] = arr[1][0] + arr[1][1] + arr[1][2] + arr[2][1] + arr[3][0] + arr[3][1] + arr[3][2]
//	arr2[1][1] = arr[1][1] + arr[1][2] + arr[1][3] + arr[2][2] + arr[3][1] + arr[3][2] + arr[3][3]
//	arr2[1][2] = arr[1][2] + arr[1][3] + arr[1][4] + arr[2][3] + arr[3][2] + arr[3][3] + arr[3][4]
//	arr2[1][3] = arr[1][3] + arr[1][4] + arr[1][5] + arr[2][4] + arr[3][3] + arr[3][4] + arr[3][5]
//
//	arr2[2][0] = arr[2][0] + arr[2][1] + arr[2][2] + arr[3][1] + arr[4][0] + arr[4][1] + arr[4][2]
//	arr2[2][1] = arr[2][1] + arr[2][2] + arr[2][3] + arr[3][2] + arr[4][1] + arr[4][2] + arr[4][3]
//	arr2[2][2] = arr[2][2] + arr[2][3] + arr[2][4] + arr[3][3] + arr[4][2] + arr[4][3] + arr[4][4]
//	arr2[2][3] = arr[2][3] + arr[2][4] + arr[2][5] + arr[3][4] + arr[4][3] + arr[4][4] + arr[4][5]
//
//	arr2[3][0] = arr[3][0] + arr[3][1] + arr[3][2] + arr[4][1] + arr[5][0] + arr[5][1] + arr[5][2]
//	arr2[3][1] = arr[3][1] + arr[3][2] + arr[3][3] + arr[4][2] + arr[5][1] + arr[5][2] + arr[5][3]
//	arr2[3][2] = arr[3][2] + arr[3][3] + arr[3][4] + arr[4][3] + arr[5][2] + arr[5][3] + arr[5][4]
//	arr2[3][3] = arr[3][3] + arr[3][4] + arr[3][5] + arr[4][4] + arr[5][3] + arr[5][4] + arr[5][5]
//
//	k := arr2[0][0]
//
//	for i := 0; i < len(arr2); i++ {
//		for j := 0; j < len(arr2); j++ {
//			if k < arr2[i][j] {
//				k = arr2[i][j]
//			}
//		}
//	}
//	fmt.Println(k)
//
// }
//func main() { //Волшебный квдарат
//	sum := [8]int32{}
//	s := [][]int32{
//		{4, 8, 2},
//		{4, 5, 7},
//		{6, 1, 6},
//	}
//
//	a1 := [3][3]int32{
//		{2, 9, 4},
//		{7, 5, 3},
//		{6, 1, 8},
//	}
//
//	a2 := [3][3]int32{
//		{2, 7, 6},
//		{9, 5, 1},
//		{4, 3, 8},
//	}
//	a3 := [3][3]int32{
//		{8, 3, 4},
//		{1, 5, 9},
//		{6, 7, 2},
//	}
//	a4 := [3][3]int32{
//		{8, 1, 6},
//		{3, 5, 7},
//		{4, 9, 2},
//	}
//	a5 := [3][3]int32{
//		{4, 9, 2},
//		{3, 5, 7},
//		{8, 1, 6},
//	}
//	a6 := [3][3]int32{
//		{4, 3, 8},
//		{9, 5, 1},
//		{2, 7, 6},
//	}
//	a7 := [3][3]int32{
//		{6, 7, 2},
//		{1, 5, 9},
//		{8, 3, 4},
//	}
//	a8 := [3][3]int32{
//		{6, 1, 8},
//		{7, 5, 3},
//		{2, 9, 4},
//	}
//
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			sum[0] = sum[0] + int32(math.Abs(float64(s[i][j]-a1[i][j])))
//			sum[1] = sum[1] + int32(math.Abs(float64(s[i][j]-a2[i][j])))
//			sum[2] = sum[2] + int32(math.Abs(float64(s[i][j]-a3[i][j])))
//			sum[3] = sum[3] + int32(math.Abs(float64(s[i][j]-a4[i][j])))
//			sum[4] = sum[4] + int32(math.Abs(float64(s[i][j]-a5[i][j])))
//			sum[5] = sum[5] + int32(math.Abs(float64(s[i][j]-a6[i][j])))
//			sum[6] = sum[6] + int32(math.Abs(float64(s[i][j]-a7[i][j])))
//			sum[7] = sum[7] + int32(math.Abs(float64(s[i][j]-a8[i][j])))
//
//		}
//	}
//	var k int32 = sum[0]
//	for i := 0; i < 8; i++ {
//		if k > sum[i] {
//			k = sum[i]
//		}
//	}
//	fmt.Println(k)
//}
