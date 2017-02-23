package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/thomasheller/slicecmp"
)

var (
	approx map[int]string
	bytes  map[int]string
	names  map[int]string
)

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

	rows := make([][]string, 0)

	if len(os.Args) == 2 {
		exp := parseInt(os.Args[1])
		row := pow2(exp)
		rows = append(rows, row)
	} else if len(os.Args) == 3 {
		exp1 := parseInt(os.Args[1])
		exp2 := parseInt(os.Args[2])
		for i := exp1; i <= exp2; i++ {
			fmt.Println("append " + string(i))
			row := pow2(i)
			rows = append(rows, row)
		}
	} else {
		for i := 1; i <= 40; i++ {
			row := pow2(i)
			rows = append(rows, row)
		}
	}

	rows = slicecmp.Transform(rows)
	fmt.Println(slicecmp.Sprintf('=', 4, slicecmp.AlignRight, []string{"2^x", "Value", "Approx.", "Mnemonic", "Byte size"}, rows...))
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

func parseInt(s string) int {
	exp, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic("Not a valid integer")
	}
	return int(exp)
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
