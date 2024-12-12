package handlers

import (
	"log"
	"net/http"
	"os" // 用於讀取環境變數

	voting "lunchking/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func GetCandidatesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 連接到以太坊客戶端（Ganache）
		client, err := ethclient.Dial("http://127.0.0.1:7545") // Ganache 節點地址
		if err != nil {
			log.Printf("Failed to connect to Ethereum client: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Ethereum client"})
			return
		}

		// 從環境變數中讀取合約地址
		contractAddress := os.Getenv("CONTRACT_ADDRESS")
		if contractAddress == "" {
			log.Println("CONTRACT_ADDRESS is not set")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Contract address is not set"})
			return
		}

		instance, err := voting.NewVoting(common.HexToAddress(contractAddress), client)
		if err != nil {
			log.Printf("Failed to load contract instance: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load contract instance"})
			return
		}

		candidates, err := instance.GetAllCandidates(nil) // 假設智能合約有這個方法
		if err != nil {
			log.Printf("Failed to get candidates: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get candidates"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"candidates": candidates})
	}
}
