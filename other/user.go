package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

const key = `{"address":"1cc10ed74aeb408ce20907827ccfad1684fd99a0","crypto":{"cipher":"aes-128-ctr","ciphertext":"50379ea081002ddbe085a7e0e5ea2d6a034af7c16ff9d658dd3de39d6d779348","cipherparams":{"iv":"bacf07b87ec942359ea17008772945a8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"63d8e789d36be58cd03f1d4e2ed3e5ebf095d70ea091603935c230ca587545d5"},"mac":"a4da7071624b731ac69856444ba7dc0e836f95211bdaab6a87f2a2e0d0280398"},"id":"696720d4-d4a5-4fb6-8d69-0d5d22b8e74c","version":3}`
const contractAddress = "0x9084b5E7a14B87635aF23f4B77C728aC906486cd"

func (token *Abi) regis(auth *bind.TransactOpts, id uint8, passwd string, _passwdAgain string) {
	_, err := token.Register(auth, id, passwd, _passwdAgain)
	if err != nil {
		log.Fatalf("transaction produce fail: %v", err)
	}
	fmt.Println("register is true")
}

func (token *Abi) logs(auth *bind.TransactOpts, id uint8, passwd string) {
	_, err := token.Login(auth, id, passwd)

	if err != nil {
		log.Fatalf("login is fail: %v", err)
	}
	fmt.Println("login is true")
}

func (token *Abi) setPass(auth *bind.TransactOpts, id uint8, passwd string, _passwdAgain string) {
	_, err := token.SetPassword(auth, id, passwd, _passwdAgain)

	if err != nil {
		log.Fatalf("reset password is fail: %v", err)
	}
}

func main() {
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	fmt.Println("connect to localgit geth node...", conn)
	if err != nil {
		log.Fatalf("could not connect to local node: %v", err)
	}
	defer conn.Close()
	token, err := NewAbi(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	fmt.Println("contract token======>:", token)

	auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
	if err != nil {
		log.Fatalf("could not create auth: %v", err)
	}

	fmt.Println("=========start put data=========")
	var choose int
	fmt.Println("选择功能:1、注册,2、登录,3、修改密码")
	fmt.Scan(&choose)
	switch choose {
	case 1:
		var id uint8
		var password string
		var _password string
		fmt.Println("请输入id")
		fmt.Scan(&id)
		fmt.Println("请输入密码")
		fmt.Scan(&password)
		fmt.Println("请再次输入密码")
		fmt.Scan(&_password)
		token.regis(auth, id, password, _password)
	case 2:
		var id uint8
		var password string
		fmt.Println("请输入id")
		fmt.Scan(&id)
		fmt.Println("请输入密码")
		fmt.Scan(&password)
		token.logs(auth, id, password)
	case 3:
		var id uint8
		var password string
		var _password string
		fmt.Println("请输入id")
		fmt.Scan(&id)
		fmt.Println("请输入原密码")
		fmt.Scan(&password)
		fmt.Println("请输入新密码")
		fmt.Scan(&_password)
		token.setPass(auth, id, password, _password)
	}

	fmt.Println("put data to smart contract success!")

}
