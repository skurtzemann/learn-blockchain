package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {
	// // int playing
	// a := big.NewInt(1)
	// fmt.Printf("%b\n", a)
	// a.Lsh(a, uint(3))
	// fmt.Printf("%b\n", a)
	// fmt.Println(a)

	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, uint64(256))
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%b\n", buff)
	fmt.Println(buff.Bytes())
	fmt.Printf("%x\n", buff.Bytes())
}

//
//If a := 1
//Lsh(1) => 10 => 2
//Lsh(3) => 1000 => 8
