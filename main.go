package main

// github.com/divan/qrlogo

// 1 (соответствует Windows-1251) – будет выводиться ST00011;
// 2 (соответствует UTF-8) – будет выводиться ST00012.

func main() {
	QRgen("http://yandex.ru", "gen/qr.png")
	QRreader("gen/qr.png")
}
