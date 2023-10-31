package handler

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type SSLCheckRequest struct {
	URL  string `json:"url"`
	PORT string `json:"port"`
}

func SSLCheck(c *fiber.Ctx) error {
	ssl := new(SSLCheckRequest)
	if err := c.BodyParser(ssl); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	if ssl.URL == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "URL must be provided")
	}
	if ssl.PORT == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "PORT must be provided")
	}

	expiry, organization, err := checkSSLCertificate(ssl.URL, ssl.PORT)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(fiber.Map{"ssl_expiry": expiry.Format(time.RFC850), "organization": organization})
}

func checkSSLCertificate(url, port string) (time.Time, string, error) {
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", url, port), nil)
	if err != nil {
		return time.Time{}, "", err
	}
	defer conn.Close()

	err = conn.VerifyHostname(url)
	if err != nil {
		return time.Time{}, "", fmt.Errorf("Hostname doesn't match with certificate: %s", err)
	}

	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	organization := conn.ConnectionState().PeerCertificates[0].Issuer.Organization[0]

	return expiry, organization, nil
}
