package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math"
	"math/big"
	"strings"
)

var client *ethclient.Client

func init() {
	client, _ = ethclient.Dial("https://ropsten.infura.io/v3/bfbcd2c178fb4f899b13c4548504664e")
}

// 获取余额
func GetBalance(wallet_addr string) *big.Int {
	account := common.HexToAddress(wallet_addr)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034
	return balance
}

func WeiToEth(balance string) (eth_value *big.Float){
	fbalance := new(big.Float)
	fbalance.SetString(balance)
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041
	return ethValue
}

func EthToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18 - len(fracStr))
	fracInt, _ :=  new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei;
}


func main() {
	//Test()
	a := WeiToEth("1173179952907945817")
	b := EthToWei(a)
	fmt.Println(a)
	fmt.Println(b)

}