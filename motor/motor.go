package motor

type Maju interface {
	Cepat()
	Lambat()
}

type Motor struct {
	Warna           string
	jmlBan, jmlGigi int
}

func (m *Motor) Cepat(kecepatan float32) float32 {
	return kecepatan * 2
}

func (m *Motor) Lambat(kecepatan float32) float32 {
	return kecepatan / 2
}

func CreateMotor(warna string, jmlban, jmlgigi int) *Motor {
	return &Motor{
		Warna:   warna,
		jmlBan:  jmlban,
		jmlGigi: jmlgigi,
	}
}
