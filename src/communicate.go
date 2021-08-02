package main

// TODO コンセンサスアルゴリズム
// 新しいブロックが掘れたら
//   同じブロックチェーンを保持するほかノードすべてに通知
// 新しいブロックを受け取ったら
//   previousHashが有効なものか確認
//
// TODO 新しいブロックが掘れた時のP2Pアルゴリズム

type Blockchain struct {
	BlockNode Block `json:"block_node"`
}

func GetTransaction(chainName string) {

}
