package main

import "fmt"

// channel
// select

func main() {
		// 複数のチャネルを動かすときに、いずれかのチャネルが受信できなくなると他のプログラムも停止してしまう
		// 失敗例
		// ch1 := make(chan int, 2)
		// var ch2 chan string = make(chan string, 2)

		// ch2 <- "A" // ch2に文字列を送る

		// var v1 int = <- ch1 // ch1は空だからv1には何も入らず、デッドロックになる
		// var v2 string = <- ch2 // ここに辿り着けない
		// fmt.Println(v1)
		// fmt.Println(v2)

		// 解決するためにselect文を使って、二つのチャネルの処理を分岐させる
		ch1 := make(chan int, 2)
		var ch2 chan string = make(chan string, 2)

		ch2 <- "A" // ch2に文字列を送る
		ch1 <- 1 // ch1に整数を送る
		ch2 <- "B" // ch2に文字列を送る
		ch1 <- 2 // ch1に整数を送る

		// select文を使って、ch1とch2のどちらから受信できるかによって処理を分岐させる
		// caseには、チャネルに対する処理を書く
		// switch文とは違って最初に成立したcaseの処理だけが実行されるわけではなく、どの処理が実行されるかはランダム
		// ランダムにしないと、片方のチャネルに偏ってしまう
		select {
		case v1 := <- ch1: // ch1から受信できる場合
				fmt.Println(v1 + 1000)
		case v2 := <- ch2: // ch2から受信できる場合
				fmt.Println(v2 + "!?")
		default: // どちらも受信できない場合
				fmt.Println("どちらも受信できません")
		}

		// selectの活用例
		// 出力結果は実行する度に変わる

		ch3 := make(chan int)
		ch4 := make(chan int)
		ch5 := make(chan int)

		// reciever
		go func() {
				for {
						i := <- ch3 // 3.ch3から受信した値をiに代入
						ch4 <- i * 2 // 4.ch4にiの2倍の値を送信
				}
		}() // 無名関数は()で即時実行できる

		go func() {
				for {
						var i2 int = <- ch4 // 5.ch4から受信した値をi2に代入
						ch5 <- i2 - 1 // 6.ch5にi2の1減らした値を送信
				}
		}()

		// main
		n := 0
		L:
		for {
				select {
				case ch3 <- n: // 1.ch3にnの値を送信
						n++ // 2.nをインクリメント
				case i3 := <- ch5: // 7.ch5から受信した値をi3に代入
						fmt.Println("recieved", i3, n) // 8.i3とnの値を出力
				// Lがないとbreakできない
				default:
						if n > 100 {
								break L // nが100を超えたら終了
						}
				}
				// 下記のように、defaultを使わずif文でbreakすることもできる
				// if n > 100 {
				// 	break // nが100を超えたら終了
				// }
		}
}
