package Math

type ComplexNumber struct {
	real float64
	im   float64
}

func (r *ComplexNumber) SetRe(number float64) {
	r.real = number
}

func (r *ComplexNumber) SetIm(number float64) {
	r.im = number
}

func (r *ComplexNumber) GetRe() float64 {
	return r.real
}

func (r *ComplexNumber) GetIm() float64 {
	return r.im
}

func (r *ComplexNumber) ConstMult(value float64) {
	r.im = r.im * value
	r.real = r.real * value
}

func SumComplexNumber(com1, com2 ComplexNumber) (result ComplexNumber) {
	result.SetRe(com1.GetRe() + com2.GetRe())
	result.SetIm(com1.GetIm() + com2.GetIm())

	return result
}

func MultComplexNumber(com1, com2 ComplexNumber) (result ComplexNumber) {
	result.SetRe(com1.GetRe()*com2.GetRe() - com1.GetIm()*com2.GetIm())
	result.SetIm(com1.GetRe()*com2.GetIm() + com1.GetIm()*com2.GetRe())

	return result
}

func DivComplexNumber(dividend, divisor ComplexNumber) (result ComplexNumber) {
	conj := divisor
	conj.SetIm(-1 * (divisor.GetIm()))
	divisor_conj := MultComplexNumber(conj, divisor)
	dividend_conj := MultComplexNumber(conj, dividend)

	divisor_conj.SetRe(1 / divisor_conj.GetRe())

	result = MultComplexNumber(dividend_conj, divisor_conj)

	return result
}
