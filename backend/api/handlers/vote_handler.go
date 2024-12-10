package handlers

import (
	"math/big"
	"net/http"

	voting "lunchking/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func VoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			CandidateID int64 `json:"candidate_id"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		client, _ := ethclient.Dial("http://127.0.0.1:7545")
		instance, _ := voting.NewVoting(common.HexToAddress("0xYourContractAddress"), client)

		privateKey, _ := crypto.HexToECDSA("YourPrivateKeyInHex")
		auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))

		tx, err := instance.Vote(auth, big.NewInt(request.CandidateID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to vote"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Vote successful", "transaction": tx.Hash().Hex()})
	}
}
