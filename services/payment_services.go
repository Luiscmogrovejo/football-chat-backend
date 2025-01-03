package services

import (
	"errors"
	"log"

	"github.com/yourusername/football-chat-backend/config"
	"github.com/yourusername/football-chat-backend/models"
)

// SetUserSubscription sets the user's plan to premium
func SetUserSubscription(userID uint, plan string) error {
	var user models.User
	if err := config.Config.DB.First(&user, userID).Error; err != nil {
		return err
	}

	user.Plan = plan
	if err := config.Config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// For Apple IAP verification, you'd typically call Apple's verifyReceipt endpoint:
func VerifyAppleReceipt(userID uint, receiptData string) error {
	// TODO: call Apple's endpoint with the receiptData
	// If valid, upgrade user to "premium"
	// This is a placeholder
	log.Println("Verifying Apple Receipt...", receiptData)

	// Simulate success:
	return SetUserSubscription(userID, "premium")
}

// For Google IAP verification, you'd call the Google Play Developer API:
func VerifyGoogleReceipt(userID uint, purchaseToken string) error {
	// TODO: call Google Play API to verify
	log.Println("Verifying Google Purchase...", purchaseToken)

	// Simulate success:
	return SetUserSubscription(userID, "premium")
}

// If using Stripe for web payment:
func ChargeWithStripe(userID uint, token string, amount int64) error {
	// Implementation with Stripe's Go SDK
	// https://stripe.com/docs/payments/accept-a-payment?lang=go
	// For now, placeholder
	return errors.New("not implemented")
}
