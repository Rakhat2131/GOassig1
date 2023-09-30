package main

import (
	"fmt"
)

type CompressionStrategy interface {
	Compress(file string)
}
type ZipCompressionStrategy struct{}

func (z ZipCompressionStrategy) Compress(file string) {
	fmt.Printf("Compressing %s using ZIP compression\n", file)
}

type RarCompressionStrategy struct{}

func (r RarCompressionStrategy) Compress(file string) {
	fmt.Printf("Compressing %s using RAR compression\n", file)
}

type Context struct {
	CompressionStrategy
}

func (c *Context) SetCompressionStrategy(strategy CompressionStrategy) {
	c.CompressionStrategy = strategy
}
func (c *Context) CompressFile(file string) {
	c.CompressionStrategy.Compress(file)
}

func main() {
	context := Context{}
	context.SetCompressionStrategy(ZipCompressionStrategy{})
	context.CompressFile("document.txt")

	context.SetCompressionStrategy(RarCompressionStrategy{})
	context.CompressFile("image.jpg")
}
