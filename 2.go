package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/thomasheller/slicecmp"
)

var (
	approx map[int]string
	bytes  map[int]string
	names  map[int]string
)

type chunk interface {
	Begin() int
	End() int
}

type scalar struct {
	i int
}

func (s *scalar) Begin() int {
	return s.i
}

func (s *scalar) End() int {
	return s.i
}

type interval struct {
	begin int
	end   int
}

func (i *interval) Begin() int {
	return i.begin
}

func (i *interval) End() int {
	return i.end
}

func main() {
	approx = map[int]string{
		7:  "1 hundred",
		10: "1 thousand",
		20: "1 million",
		30: "1 billion",
		40: "1 trillion",
	}

	bytes = map[int]string{
		10: "1 KB",
		16: "64 KB",
		20: "1 MB",
		30: "1 GB",
		32: "4 GB",
		40: "1 TB",
	}

	names = map[int]string{
		0: "B",
		1: "KB",
		2: "MB",
		3: "GB",
		4: "TB",
		5: "PB",
		6: "EB",
		7: "ZB",
		8: "YB",
	}

	chunks := make([]chunk, 0)

	re := regexp.MustCompile(`^([\d]+)-([\d]+)$`)

	if len(os.Args) == 2 {
		for _, s := range strings.Split(os.Args[1], ",") {
			exp, err := parseInt(s)
			if err == nil {
				chunks = append(chunks, &scalar{i: exp})
				continue
			}
			if re.MatchString(s) {
				m := re.FindStringSubmatch(s)
				i1, _ := parseInt(m[1])
				i2, _ := parseInt(m[2])
				chunks = append(chunks, &interval{begin: i1, end: i2})
				continue
			}
			panic("Can't parse " + s)
		}
	} else {
		chunks = append(chunks, &interval{begin: 1, end: 40})
	}

	rows := make([][]string, 0)

	for _, chunk := range chunks {
		if chunk.Begin() <= chunk.End() {
			for i := chunk.Begin(); i <= chunk.End(); i++ {
				rows = append(rows, pow2(i))
			}
		} else {
			for i := chunk.Begin(); i >= chunk.End(); i-- {
				rows = append(rows, pow2(i))
			}
		}
	}

	rows = slicecmp.Transform(rows)

	headings := []string{"2^x", "Value", "Approx.", "Mnemonic", "Byte size"}
	fmt.Println(slicecmp.Sprintf('=', 4, slicecmp.AlignRight, headings, rows...))
}

func pow2(exponent int) []string {
	value := math.Pow(2, float64(exponent))
	return []string{
		fmt.Sprintf("%d", exponent),
		humanizeIntFloat64(exponent, value),
		approx[exponent],
		bytes[exponent],
		byteSize(exponent, value),
	}
}

func byteSize(exponent int, value float64) string {
	return fmt.Sprintf("%.0f %s", value/math.Pow(2, float64(exponent/10)*10), getName(exponent))
}

func parseInt(s string) (int, error) {
	exp, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(exp), nil
}

func humanizeIntFloat64(exponent int, value float64) (humanized string) {
	if exponent < 62 {
		humanized = humanize.Comma(int64(value))
	} else {
		humanized = fmt.Sprintf("%.0f", value)
	}
	return
}

func getName(exponent int) (name string) {
	if exponent < len(names)*10 {
		name = names[exponent/10]
	} else {
		name = fmt.Sprintf("%s+%d", names[len(names)-1], exponent/10-len(names)+1)
	}
	return
}
