package imap

import (
	"crypto/tls"
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func login(conn *tls.Conn, login string, pass string) []string {
	prefix := "a0001"
	fmt.Println("Client->", Green(prefix+" login "+login+"  ******"))
	fmt.Fprintf(conn, prefix+" login "+login+" "+pass+"\n")

	return readBeforPrefixLine(conn, prefix)
}

func examine(conn *tls.Conn) []string {
	prefix := "a0002"
	fmt.Println("Client->", Green(prefix+" examine inbox"))
	fmt.Fprintf(conn, prefix+" examine inbox\n")

	return readBeforPrefixLine(conn, prefix)
}

func fetchHeader(conn *tls.Conn) []string {
	prefix := "a0003"
	fmt.Println("Client->", Green(prefix+" fetch 1:* (ENVELOPE) "))
	fmt.Fprintf(conn, prefix+" fetch 1:* (ENVELOPE) \n")

	return readBeforPrefixLine(conn, prefix)
}

func searchUnseen(conn *tls.Conn) []string {
	prefix := "a0004"
	fmt.Println("Client->", Green(prefix+" search unseen"))
	fmt.Fprintf(conn, prefix+" search unseen\n")

	return readBeforPrefixLine(conn, prefix)
}

func logout(conn *tls.Conn) []string {
	prefix := "a0006"
	fmt.Println("Client->", Green(prefix+" logout"))
	fmt.Fprintf(conn, prefix+" logout\n")

	return readBeforPrefixLine(conn, prefix)
}