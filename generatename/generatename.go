package generatename

import (
	"time"
)

// currentMonth returns the current month number.
func currentMonth() int {
	return int(time.Now().Month())
}

// currentYear returns the current year in 06 format.
func currentYear() string {
	return time.Now().Format("06")
}

// formatedDatePTBR returns the current month and year in Jan06 format in Pt-BR.
func formatedDatePTBR() string {
	var monthsPTBR = map[int]string{
		1:  "Jan",
		2:  "Fev",
		3:  "Mar",
		4:  "Abr",
		5:  "Mai",
		6:  "Jun",
		7:  "Jul",
		8:  "Ago",
		9:  "Set",
		10: "Out",
		11: "Nov",
		12: "Dez",
	}
	return monthsPTBR[currentMonth()] + currentYear()

}

// ExpectedCSVName returns the expected csv filename at the current time.
func ExpectedCSVName() string {
	return formatedDatePTBR() + " - Veículos HabilitadosCSV"
}

func ExpectedJsonName() string {
	return formatedDatePTBR() + " - Veículos HabilitadosJSON"
}
