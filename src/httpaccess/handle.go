package httpaccess

// 参考自 https://studygolang.com/articles/9467

import (
	"block"
	"blockchain"
	"encoding/json"
	"fmt"
	"net/http"
	"transaction"
)

//HTTPAccess : 通过HTTP连接操作、查询区块链
type HTTPAccess struct {
	bc *blockchain.Blockchain
}

func (access *HTTPAccess) addBlock(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	to := vars.Get("to")
	fmt.Printf("recieve request add block, with param 'to' = '%s'", to)
	if err := access.bc.AppendNewBlock([]*transaction.Tx{transaction.NewCoinbaseTx(to, "")}); err != nil {
		w.Write([]byte(fmt.Sprintf("error occured ,err = %s", err.Error())))
		return
	}
	w.Write([]byte("block add successfully"))
}

func (access *HTTPAccess) printChainData(w http.ResponseWriter, r *http.Request) {
	bcit := access.bc.NewIterator()
	var res []*block.Block
	for b := bcit.Next(); b != nil; b = bcit.Next() {
		res = append(res, b)
	}
	j, _ := json.Marshal(res)
	w.Write(j)
}

//Run : 执行HTTP操作
func (access *HTTPAccess) Run() {

	access.bc = blockchain.NewBlockchain()

	http.HandleFunc("/add", access.addBlock)
	http.HandleFunc("/query", access.printChainData)

	fmt.Println("server linstening on 0.0.0.0:40000")
	http.ListenAndServe("0.0.0.0:40000", nil)
}
