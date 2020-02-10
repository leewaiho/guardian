package main

import (
	"flag"
	"fmt"
	"github.com/LeeWaiHo/guardian/pkg/conf"
	"github.com/LeeWaiHo/guardian/pkg/service"
	"log"
	"os"
)

var _FlagNetworkID = flag.Int64("networkID", 1, "")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	log.SetPrefix("[guardian] ")
	conf.New()
}

func main() {
	createNetworkConfiguration(*_FlagNetworkID)
}

func createNetworkConfiguration(networkID int64) {
	networkNodes, err := service.New().ScanNetworkNodes(networkID)
	if err != nil {
		log.Fatalf("获取网络节点异常 error(%v)", err)
	}

	dirname := fmt.Sprintf("./out/network_%d", 1)
	if e := ensureDirExist(dirname); e != nil {
		log.Fatal(e)
	}

	gatewayFile := dirname + "/gateway.conf"
	log.Println("生成网关配置文件 ", gatewayFile)
	if e := saveToFile(gatewayFile, []byte(networkNodes.GatewayConfiguration())); e != nil {
		log.Fatal(e)
	}
	log.Printf("生成网关配置文件完成\n\n")

	log.Println("开始生成普通网络节点")
	for i, node := range networkNodes.ClientNodes {
		clientFile := fmt.Sprintf(dirname+"/%s.conf", node.Hostname)
		log.Printf("[%d/%d]生成客户端配置文件: %s\n", i+1, len(networkNodes.ClientNodes), clientFile)
		if e := saveToFile(clientFile, []byte(networkNodes.ClientConfiguration(node))); e != nil {
			log.Fatal(e)
		}
	}
}

func ensureDirExist(dirname string) error {
	if isDir(dirname) {
		return nil
	}
	return os.MkdirAll(dirname, 0755)
}

func isDir(dirname string) bool {
	fileInfo, err := os.Stat(dirname)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func saveToFile(filename string, content []byte) error {
	f, e := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		return e
	}
	defer f.Close()
	f.Write(content)
	return nil
}
