/*
* https://github.com/Marie673
* 想定環境:Windows10
* build方法:
* go build main.go block.go blockchain.go transaction.go PoW.go
* 実行方法:
* main.exe
 */
package main

import (
	"math/rand"
	"time"
)

// Proof-of-Workの難易度
// 今回のアルゴリズムの場合，数字が小さくなるほど難しい
// 最大値255,最小値1
const diff = 235

var consensusChain = &Blockchain{
	author: "Consensus Chain",
}

const nodeNUM = 5

var names = [nodeNUM]string{"okazaki", "yude", "yamakasu", "str4w", "mimimi"}
var nodes [nodeNUM]*Blockchain

func InitialNodes() {
	for i := 0; i < nodeNUM; i++ {
		nodes[i] = CreateInitialBlock(names[i])
		// nodes[i].PrintBlockchain()
	}
}

func Mining(index int) {
	nodes[index].AddBlock()
	nodes[index].block = Consensus()
}

// TO DO
// コンセンサスされたノードのファイル出力
// フレームワークを用いてネット上でのマイニング
// トランザクションを追加する処理
// コンセンサスにより捨てられるトランザクションをどうするか
// →そのまま捨てるorトランザクションが正しくブロックに追加されたかの処理を記述
func main() {
	InitialNodes()
	go func() {
		for {
			Mining(0)
		}
	}()
	go func() {
		for {
			Mining(1)
		}
	}()
	go func() {
		for {
			Mining(2)
		}
	}()
	go func() {
		for {
			Mining(3)
		}
	}()
	go func() {
		for {
			Mining(4)
		}
	}()

	// トランザクション追加のテスト処理
	go func() {
		t1 := Transaction{
			sender:    "test1",
			recipient: "test1",
		}
		t2 := Transaction{
			sender:    "test2",
			recipient: "test2",
		}
		for {
			if rand.Int()%2 == 0 {
				t1.amount = rand.Int()
				AddTransaction(t1)
			} else {
				t2.amount = rand.Int()
				AddTransaction(t2)
			}
			randTime := time.Duration(rand.Int() % 5)
			time.Sleep(randTime * time.Second)
		}
	}()

	/*
		// 5秒毎に各ノードの状態を確認
		go func() {
			for {
				time.Sleep(5 * time.Second)
				for i := 0; i < nodeNUM; i++ {
					nodes[i].PrintBlockchain()
				}
			}
		}()
	*/

	// 5分間実行
	time.Sleep(300 * time.Second)

	// TO DO
	// 各ノードの状態をファイルで出力
	/*
		file, err := os.Create(`result\\mining.txt`)
		if err != nil {
			panic(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		ls := Consensus()
		for i := 0; i < len(ls); i++ {
			m, err := json.Marshal(ls[i])
			if err != nil {
				panic(err)
			}
			output := string(m)
			output += "\r\n"
			_, err = file.Write(([]byte)(output))
			if err != nil {
				return
			}
		}
	*/
}
