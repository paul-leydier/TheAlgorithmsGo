package pi

import (
	"fmt"
	"math"
	"math/big"
)

const prec = 15000

// https://en.wikipedia.org/wiki/Chudnovsky_algorithm
func Chudnovsky(q int) (*big.Float, int) {
	digitsPerRound := math.Log10(151931373056000)
	C := new(big.Float).SetPrec(prec).Mul(big.NewFloat(426880), new(big.Float).Sqrt(big.NewFloat(10005)))
	mChan, lChan, xChan := make(chan big.Float), make(chan big.Float), make(chan big.Float)
	go multinomialTerm(q, mChan)
	go linearTerm(q, lChan)
	go exponentialTerm(q, xChan)
	termsSum, tmp, pi := new(big.Float), new(big.Float), new(big.Float)
	termsSum.SetPrec(prec)
	tmp.SetPrec(prec)
	pi.SetPrec(prec)
	for i := 0; i <= q; i++ {
		m := <-mChan
		l := <-lChan
		x := <-xChan
		tmp.Mul(&m, &l)
		tmp.Quo(tmp, &x)
		termsSum.Add(termsSum, tmp)
		pi.Quo(C, termsSum)
		fmt.Println(pi.Text('f', int(digitsPerRound*float64(i))))
	}

	return pi, int(digitsPerRound * float64(q))
}

func multinomialTerm(q int, c chan<- big.Float) {
	prevM := big.NewFloat(1).SetPrec(prec)
	c <- *prevM
	for i := 1.0; i <= float64(q); i++ {
		tmp := (12*i - 2) * (12*i - 6) * (12*i - 10)
		prevM.Mul(prevM, big.NewFloat(tmp))
		prevM.Quo(prevM, big.NewFloat(i*i*i))
		c <- *prevM
	}
}

func linearTerm(q int, c chan<- big.Float) {
	prevL := big.NewFloat(13591409).SetPrec(prec)
	c <- *prevL
	for i := 1; i <= q; i++ {
		prevL.Add(prevL, big.NewFloat(545140134))
		c <- *prevL
	}
}

func exponentialTerm(q int, c chan<- big.Float) {
	prevX := big.NewFloat(1).SetPrec(prec)
	c <- *prevX
	for i := 1; i <= q; i++ {
		prevX.Mul(prevX, big.NewFloat(-262537412640768000))
		c <- *prevX
	}
}
