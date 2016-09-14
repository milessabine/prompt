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

func (p Prompt) Scan(n string) string {
	fmt.Print(n, " ")
	p.scanner.Scan()
	return p.scanner.Text()
}

func (p Prompt) ScanInt(n string) (int, error) {
	s := p.Scan(n)
	return strconv.Atoi(s)
	/*i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil*/
}

func (p Prompt) ScanFloat64(n string) (float64, error) {
	s := p.Scan(n)
	return strconv.ParseFloat(s, 64)
	/*f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0,err
	}
	return f, nil*/
}
