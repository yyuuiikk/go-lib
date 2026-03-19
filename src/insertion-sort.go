package main

import (
	"fmt"
)

// 挿入ソート
func main() {
	arr1 := []int{3, 2, 9, 1, 5, 6}
	for j := 1; j < len(arr1); j++ {
		key := arr1[j]
		i := j - 1

		// fmt.Println("key:", key)
		//fmt.Fprintln(os.Stdout, "i:", i)

		for i >= 0 && arr1[i] > key {
			// 大きい値であれば、右にずらす（代入して書き換える）
			arr1[i+1] = arr1[i]

			// fmt.Println(arr1)
			i--
		}
		// 比較対象の値を最後に移動させる
		arr1[i+1] = key
		// fmt.Println("temp_sorted:", arr1)
	}
	fmt.Println(arr1)
}
