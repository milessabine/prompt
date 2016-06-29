package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Prompt struct {
	scanner *bufio.Scanner
}

func New() Prompt {
	s := bufio.NewScanner(os.Stdin)
	return Prompt{
		scanner: s,
	}
}

func (p *Prompt) Scan(n string) string {
	fmt.Print(n, " ")
	p.scanner.Scan()
	return p.scanner.Text()
}

func (p *Prompt) ScanInt(n string) int {
	s := p.Scan(n)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Not a valid number", s, err)
		return 0
	}
	return i
}

func (p *Prompt) ScanFloat64(n string) float64 {
	s := p.Scan(n)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Please input valid number", s, err)
		return 0.0
	}
	return f
}
