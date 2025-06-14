package enums

type IrrigationType string

const (
	ASPERSAO    IrrigationType = "ASPERSAO"
	GOTEJAMENTO IrrigationType = "GOTEJAMENTO"
)

func IsValidateFieldIrrigationTypeEnum(irrigationType IrrigationType) bool {

	switch irrigationType {
	case ASPERSAO, GOTEJAMENTO:
		return true
	}

	return false

}
