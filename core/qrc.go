// Package qrc defines the core functionalities of QR code generation
package qrc

import (
	"net/url"

	qrcode "github.com/skip2/go-qrcode"
)

// GenerateQRCodeFromURLString generate QR from url string
func GenerateQRCodeFromURLString(urlstring, outputfilename string) error {
	_, err := url.ParseRequestURI(urlstring)
	if err != nil {
		return err
	}

	return qrcode.WriteFile(urlstring, qrcode.Highest, 256, outputfilename)
}
