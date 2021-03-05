package main
 
import (
	"fmt"
	"io"
	"os"
)
 
func main() {
 
	// Простая запись строки
	io.WriteString(os.Stdout,
		"Это строка для стандартного вывода.\n")
 
	io.WriteString(os.Stderr,
		"Это строка для стандартного вывода ошибки.\n")
 
	// Stdout/err имплементации
	// интерфейс записи
	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 200; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}
 
	// Пакет fmt также можно использовать
	fmt.Fprintln(os.Stdout, )
}