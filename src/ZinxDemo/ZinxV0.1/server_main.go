package main

import "Zinx/znet"

func main() {
	s := znet.NewServer("tcp4", "0.0.0.0", 7000, "渣渣辉传奇")
	s.Server()
}
