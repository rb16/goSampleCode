package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)                                    // 5, which is n(hello)
	c = 0                                             // reset the counter
	fmt.Println(fmt.Fprintf(&c, "hello, %s", "rana")) //Fprintf returns no. bytes to written and error
	fmt.Println(c)                                    // 11, which is len of "hello, rana"
}
