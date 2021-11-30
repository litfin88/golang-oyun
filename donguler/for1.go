package donguler

import (
	"fmt"
	"math/rand"
	"time"
)

func Demo1() {
	rand.Seed(time.Now().UnixNano())
	aklimdakiSayi := rand.Intn(100)
	tahmin := 0

	tahminSayisi := 0

	for aklimdakiSayi != tahmin {
		fmt.Print("Sayı tahmin edin: ")
		fmt.Scanln(&tahmin)

		if tahmin < aklimdakiSayi {
			fmt.Println("Yukarı!")
			tahminSayisi++
		} else if tahmin > aklimdakiSayi {
			fmt.Println("Aşağı!")
			tahminSayisi++
		}
	}

	if tahmin == aklimdakiSayi {
		fmt.Println("Bildin!")
	}

	if tahminSayisi <= 2 {
		fmt.Println("Süper!")
	} else if tahminSayisi <= 5 {
		fmt.Println("Güzel!")
	} else if tahminSayisi <= 8 {
		fmt.Println("Fena Değil!")
	} else {
		fmt.Println("Çok kötü.")
	}
}
