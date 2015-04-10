package transaction

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}


func TestIncome(t *testing.T) {
	var series []float64
	// Regular Income, $7000, starting in 1st October 
	
//	RegularIncome(7000, "USD", "01/10/")

}