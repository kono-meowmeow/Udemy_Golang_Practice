package foo
// mainパッケージから分割されたパッケージとなる
// fooパッケージに定義したものを、mainパッケージで呼び出す

const (
		// 大文字で始まっているので、他のパッケージからも参照可能
		Max = 100
		// 小文字で始まっているので、他のパッケージからは参照できない
		min = 1
)

// 大文字で始まっているので、他のパッケージからも参照可能
func ReturnMin() int {
		// 同じパッケージ内なので、minを参照できる
		return min
}
