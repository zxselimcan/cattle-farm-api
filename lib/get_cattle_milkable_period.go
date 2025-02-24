package lib

import (
	"api/models"
	"time"
)

func GetCattleMilkablePeriod(cattle *models.Cattle) string {
	ONE_DAY := (time.Hour * 24)
	_305_DAYS_AGO := time.Now().Add(-305 * ONE_DAY)
	_250_DAYS_AGO := time.Now().Add(-250 * ONE_DAY)
	_210_DAYS_AGO := time.Now().Add(-210 * ONE_DAY)

	if !cattle.IsAlive {
		return "DEAD"
	}

	if cattle.Gender != "female" {
		return "MALE"
	}

	if cattle.PregnancyStatus == "pregnant" || cattle.PregnancyStatus == "inseminated" {
		// hamileyse ve tohumlama üzerinden 250 gün geçmişse
		if cattle.LastInseminationDate.Before(_250_DAYS_AGO) {
			return "DRY_2"
		}
		//hamileyse ve tohumlama üzerinden 210 gün geçmişse
		if cattle.LastInseminationDate.Before(_210_DAYS_AGO) {
			return "DRY_1"
		}
		// hamileyse ve tohumlama üzerinden 210 gün geçmemişse

		if cattle.ChildrenCount < 1 {
			return "PREGNANT"
		}

		return "MILKABLE"
	}

	// hamile değilse ve hiç doğum yapmamışsa
	if cattle.ChildrenCount < 1 {
		return "NOT_READY"
	}
	// hamile değilse ve doğum üzerinden 300 gün geçmişse
	if cattle.LastGiveBirthDate.Before(_305_DAYS_AGO) {
		return "DRY_3"
	}
	// hamile değilse ve doğum üzerinden 300 gün geçmemişse
	return "MILKABLE"

}
