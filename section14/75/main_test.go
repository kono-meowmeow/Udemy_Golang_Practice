package main

import (
		// テストコードを書くには、testingパッケージをインポートする
		"testing"
)

// Debugの値をtrueにすると、テストをスキップする
var Debug bool = false

// テスト用の処理を書いていく
// テスト用の関数は、Testから始まる名前とする
// Testの後に、テストしたい関数名を書く
// テスト用の関数は、引数に*testing.Tを取る

func TestIsOne(t *testing.T) {
		var i int = 1 // iが1でなければ、テストは失敗する
		// Debugがtrueの場合は、このテストは行わない
		if Debug {
				t.Skip("スキップしている")
		}

		// mainパッケージのテストを行う
		// テストしたい関数(mainパッケージのIsOne()という関数)を呼び出す
		v := IsOne(i) // vには、IsOne()の戻り値(bool値)が入る

		// vがtrueであれば、テストは成功
		// vがfalseであれば、テストは失敗で、エラーが出力される
		if !v {
				t.Errorf("%v != %v", i, 1)
		}
}

// テストを実行するには、go testを実行する(ターミナルで実行する)
// go testを実行すると、テストが成功したかどうかが表示される
// テストが成功した場合は、okと表示される
// テストが失敗した場合は、FAILと表示され、エラー内容が表示される。ここでのエラー内容は、t.Errorf()で出力した内容

// go test -v を実行すると、テストの詳細が表示される

// go test ./... でディレクトリ以下のテストを全て実行する
// 今回だと、mainパッケージとalibパッケージのテストが実行される
// go test -v ./... でテストの詳細が表示される

// go test -cover ./... でディレクトリ以下のテストのカバレッジを表示する
// 全ての関数のテストを行なっていると、100%と表示される
