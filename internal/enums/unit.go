package enums

type Unit string

const (
	m        Unit = "m"
	ha       Unit = "ha"
	kg       Unit = "kg"
	m2       Unit = "m2"
	m3       Unit = "m3"
	ml       Unit = "ml"
	mm       Unit = "mm"
	cm       Unit = "cm"
	saca     Unit = "saca"
	mmdia    Unit = "mm/dia"
	mmsemana Unit = "mm/semana"
	kgha     Unit = "kg/ha"
)

func IsValidateFieldUnitEnum(unit Unit) bool {

	switch unit {
	case m, ha, kg, m2, m3, ml, mm, cm, saca, mmdia, mmsemana, kgha:
		return true

	default:
		return false
	}
}
