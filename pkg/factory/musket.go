package factory

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
