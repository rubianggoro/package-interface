package sepeda

type Maju interface {
	Cepat()
	Lambat()
}

type Sepeda struct {
	Jenis string
	Roda  int
}

func (s *Sepeda) Cepat(kecepatan float32) float32 {
	return kecepatan * 2
}

func (s *Sepeda) Lambat(kecepatan float32) float32 {
	return kecepatan / 2
}

func CreateSepeda(jenis string, roda int) *Sepeda {
	return &Sepeda{
		Jenis: jenis,
		Roda:  roda,
	}
}
