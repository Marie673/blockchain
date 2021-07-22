package main

import (
	"bufio"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

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

func httpServer() {
	var e = echo.New()
	InitRouting(e)
	e.Logger.Fatal(e.Start(":10000"))
}

func InitRouting(e *echo.Echo) {
	e.GET("/", hello)
	e.GET("/blockchain/:chain_name", GetChain)
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

func GetChain(c echo.Context) error {
	chainName := c.Param("chain_name")
	fileName := "Blockchain/" + chainName

	file, err := os.Open(fileName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't open file: "+chainName)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// TODO JSON形式に変える
	var blockStr string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {

			continue
		}
		blockStr += line
	}

	return c.JSON(http.StatusOK, blockStr)
}
