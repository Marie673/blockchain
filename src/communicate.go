package main

import (
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"os"
)

// TODO コンセンサスアルゴリズム
// TODO 新しいブロックが掘れた時のP2Pアルゴリズム

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
	var chain string
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
	buf, err := ioutil.ReadAll(file)
	chain = string(buf)
	fmt.Printf("%s\n", chain)

	return c.JSON(http.StatusOK, chain)
}
