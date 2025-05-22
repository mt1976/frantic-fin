package main

import (
	"time"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-fin/banking/calendarmath"
	"github.com/mt1976/frantic-fin/financial"
)

func main() {
	examples := []string{
		"1k", "1m", "1.5m", "2b500m", "5.5m250k", "1b2m3k", "1T", "1.2t3b4.5m", // valid
		"", "foo", "1mfoo", "m1k", "10x", "1.2.3m", // invalid
	}

	for _, ex := range examples {
		val, err := financial.ParseFinancialAbbreviationToInt(ex)
		if err != nil {
			logHandler.ErrorLogger.Printf("❌ '%s' → Error: %v\n", ex, err)
		} else {
			logHandler.InfoLogger.Printf("✅ '%s' → %d (int)\n", ex, val)
		}

		val2, err := financial.ParseFinancialAbbreviationToFloat(ex)
		if err != nil {
			logHandler.ErrorLogger.Printf("❌ '%s' → Error: %v\n", ex, err)
		} else {
			logHandler.InfoLogger.Printf("✅ '%s' → %f (float)\n", ex, val2)
		}

		val3, err := financial.ParseFinancialAbbreviationToString(ex)
		if err != nil {
			logHandler.ErrorLogger.Printf("❌ '%s' → Error: %v\n", ex, err)
		} else {
			logHandler.InfoLogger.Printf("✅ '%s' → %s (string)\n", ex, val3)
		}

		val4, err := financial.ParseFinancialAbbreviationToAmountString(ex)
		if err != nil {
			logHandler.ErrorLogger.Printf("❌ '%s' → Error: %v\n", ex, err)
		} else {
			logHandler.InfoLogger.Printf("✅ '%s' → %s (amount string)\n", ex, val4)
		}
	}

	for i := 0; i < 15; i++ {
		pivot := time.Now().Add(time.Duration(i) * (time.Hour * 24))
		logHandler.InfoLogger.Printf("Pivot: %s\n", pivot.Format(time.DateOnly))
		NextBusinessDay, _ := calendarmath.GetNextWorkingDay(pivot)
		logHandler.InfoLogger.Printf("Next Business Day: %s\n", NextBusinessDay.Format(time.DateOnly))
		PreviousBusinessDay, _ := calendarmath.GetPreviousWorkingDay(pivot)
		logHandler.InfoLogger.Printf("Previous Business Day: %s\n", PreviousBusinessDay.Format(time.DateOnly))
		IsWorkingDay, _ := calendarmath.IsWorkingDay(pivot)
		logHandler.InfoLogger.Printf("Is Working Day: %t\n", IsWorkingDay)
		logHandler.InfoLogger.Printf("--------------------------------------------------\n")
	}

}
