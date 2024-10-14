package main

import (
	"fmt"
	"time"

	"github.com/rajeshreddy70133/go-licenser"
)

func generateLicense(userID string) (string, error) {
	license := licenser.License{
		UserID:      userID,
		Expires:     time.Now().Add(30 * 24 * time.Hour), // 30 days trial
		LicenseType: "trial",
	}
	key, err := licenser.GenerateLicenseKey(license, "your-secret-key")
	if err != nil {
		return "", err
	}
	return key, nil
}

func validateLicense(licenseKey string) (bool, error) {
	license, err := licenser.ValidateLicenseKey(licenseKey, "your-secret-key")
	if err != nil {
		return false, err
	}
	if license.Expires.Before(time.Now()) {
		return false, nil // License has expired
	}
	return true, nil
}

func main() {
	userID := "user123"
	// Generate license
	licenseKey, err := generateLicense(userID)
	if err != nil {
		fmt.Println("Error generating license:", err)
		return
	}
	fmt.Println("Generated License Key:", licenseKey)

	// Validate license
	valid, err := validateLicense(licenseKey)
	if err != nil {
		fmt.Println("Error validating license:", err)
		return
	}
	if valid {
		fmt.Println("License is valid.")
	} else {
		fmt.Println("License is invalid or expired.")
	}
}
