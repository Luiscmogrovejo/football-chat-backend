package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/football-chat-backend/services"
)

// SubscribeUser - starts a subscription or checks user plan
func SubscribeUser(c *gin.Context) {
	userID := c.GetUint("user_id")

	// For example, call Stripe or mark user as "premium"
	// This is a placeholder
	err := services.SetUserSubscription(userID, "premium")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to subscribe user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subscription updated"})
}

// VerifyAppleIAP - verify Apple In-App Purchase receipts
func VerifyAppleIAP(c *gin.Context) {
	var payload struct {
		ReceiptData string `json:"receipt_data" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")
	err := services.VerifyAppleReceipt(userID, payload.ReceiptData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Apple IAP verified, subscription updated"})
}

// VerifyGoogleIAP - verify Google Play purchase
func VerifyGoogleIAP(c *gin.Context) {
	var payload struct {
		PurchaseToken string `json:"purchase_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")
	err := services.VerifyGoogleReceipt(userID, payload.PurchaseToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Google IAP verified, subscription updated"})
}
