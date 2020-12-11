package taxgo 

import(
	"fmt"
	"math"
)

type Taxpayer struct {
	Name   string
	Income float64
  Year int
	Status string
  States []string
}

type Bracket struct {
	Status    string
	Threshold []float64
}

type Taxes interface {
	fedTax() float64
}

func Calc(t Taxes) {
	fmt.Println(t.fedTax())
}

func (tp Taxpayer) FedTax() float64 {
	r := make(map[int][]float64)
	r[2020] = []float64{.1, .12, .22, .24, .32, .35, .37}
	br := make(map[int][]Bracket)
	br[2020] = []Bracket{
		{"IND", []float64{0, 9875, 40125, 85525, 163300, 207350, 518400}},
		{"MFJ", []float64{0, 19750, 80250, 171050, 326600, 414700, 622050}},
	}
	var tax float64
	var myBrackets []float64

	for _, v := range br[tp.Year] {
		if tp.Status == v.Status {
			myBrackets = append(myBrackets, v.Threshold...)
		}
	}

		for i := 0; i < len(myBrackets)-1; i++ {
			if tp.Income > myBrackets[i] {
				tax += (math.Min(tp.Income, myBrackets[i+1])-myBrackets[i]) * r[tp.Year][i]
			}
		}
    if tp.Income > myBrackets[len(myBrackets)-1]{
      tax += (tp.Income - myBrackets[len(myBrackets)-1]) * r[tp.Year][len(r[tp.Year])-1]
    }

	return tax
}