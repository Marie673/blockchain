package main

// Proof-of-Workの難易度
// 今回のアルゴリズムの場合，数字が小さくなるほど難しい
// 最大値255,最小値1
const diff = 235

var limit = makeBytePoW(diff)

func main() {
	go func() {
		httpServer()
	}()
	Mining("test", "10.1.1.1")
}
