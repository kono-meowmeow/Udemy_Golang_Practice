package main

import "fmt"

// 配列型
// 後から要素数の変更ができない → 後から要素数を変更したいときはスライスを使う

func main() {
		// 配列の宣言
		var arr1 [3]int // 要素数が3のint型の配列、明示的な宣言で初期値なし
		fmt.Println(arr1)
		fmt.Printf("%T\n", arr1) // [3]int 要素数を含めて型となるから、要素数を変えられない

		var arr2 [3]string = [3]string{"A", "B"} // 3つ目の要素は空文字になる、明示的な宣言で初期値あり
		fmt.Println(arr2)

		arr3 := [3]int{1, 2, 3} // 暗黙的な宣言で初期値あり
		fmt.Println(arr3)

		arr4 := [...]string{"C", "D"} // 要素数を明示しない宣言
		fmt.Println(arr4)
		fmt.Printf("%T\n", arr4) // [2]string

		// 配列の値の取得
		fmt.Println(arr2[0]) // arr2の最初の値を取れる、[0]はインデックス番号
		fmt.Println(arr2[1])
		fmt.Println(arr2[2])
		fmt.Println(arr2[2-1]) // arr2[1]と同じ
		fmt.Println(arr2[2-2])
		fmt.Println(arr3[3-1]) // [要素数-1]にすると、配列の最後の要素を取得できる
}
