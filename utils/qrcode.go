package utils

import "github.com/skip2/go-qrcode"

func GenQrCode(url string) {
	qrcode.WriteFile(url, qrcode.Medium, 256, "/tmp/foo.png")
}
