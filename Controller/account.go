package controller

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"runtime"

	"github.com/gin-gonic/gin"
	models "github.com/lethanhsonthd/EOS_API_create_account/models"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func CreateAccount(c *gin.Context) {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	}
	// User truyen vao wallet name va wallet password, neu khong truyen vao thi su dung Wallet mac dinh (sonlt)
	defaultWalletName := "sonlt"
	defaultWalletPassword := "PW5JbTcLSUDANfYWQBAUgRFZNgK5uFUGFiraPAwNVGKH8QPnv6VLN"
	var data models.Account
	c.Bind(&data)
	// Lay thong tin tu POST request
	accountName, walletName, walletPassword := data.AccountName, data.WalletName, data.WalletPassword
	if walletName == "" {
		walletName = defaultWalletName
		walletPassword = defaultWalletPassword
	}
	usr, err := user.Current()
	check(err)
	// InfoCreateAccount: Noi chua thong tin account can tao
	InfoCreateAccount := usr.HomeDir + "/eos/build/programs/cleos/information_need_for_create_account.txt"
	file, err := os.OpenFile(InfoCreateAccount, os.O_WRONLY, 0600)
	defer file.Close()
	check(err)
	// Ghi vao file keyPair.txt cac thong tin: account name can tao, wallet name va password
	message := "ACCOUNT_NAME: " + accountName + "\n"
	message += "WALLET_NAME: " + walletName + "\n"
	message += "WALLET_PASSWORD: " + walletPassword + "\n"
	if _, err = file.WriteString(message); err != nil {
		panic(err)
	}
	fmt.Println(message)
	// Run file script tao account
	pathToScriptCreateAccount := usr.HomeDir + "/eos/build/programs/cleos/myScripts/createAccount.sh"
	out, err := exec.Command("/bin/bash", pathToScriptCreateAccount).Output()
	check(err)
	output := string(out[:])
	fmt.Println(output)
	c.AsciiJSON(http.StatusForbidden, gin.H{
		message: output,
	})
	return
}
