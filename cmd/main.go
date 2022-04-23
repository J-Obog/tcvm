package main

import "github.com/J-Obog/tcvm/pkg/vmachine"

func main() {
	vm := vmachine.VM{}

	vm.LoadFromFile("prog.obj")
	vm.Run()
}