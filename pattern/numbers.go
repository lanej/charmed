package pattern

import "github.com/lanej/charmed/text"

func (m *CohortComparison) signedNumber(index int, value float64, precision int, unit string, available bool) text.SignedNumber {
	return text.SignedNumber{Value: value, Available: available, Precision: precision, Unit: unit, Tones: m.metricSigns[m.data.Metrics[index].ID]}
}
