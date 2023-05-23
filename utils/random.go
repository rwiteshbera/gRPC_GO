package utils

import (
	"math/rand"
	"rpcProject/pb"

	"github.com/google/uuid"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	case 3:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_UNKNOWN
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomGPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}

	return randomStringFromSet(
		"Ryzen 7 Pro 2700U",
		"Ryzen 5 Pro 3500U",
		"Ryzen 3 Pro 3200GE",
	)
}

func randomGPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Dell", "Lenovo", "HP", "Acer", "Asus", "MSI")
}

func randomLaptopName(name string) string {
	switch name {
	case "Dell":
		return randomStringFromSet("Vostro", "XPS", "Alienware", "Inspiron")
	case "Lenovo":
		return randomStringFromSet("Thinkpand", "Ideapad", "Yoga", "Legion", "Notebook")
	case "HP":
		return randomStringFromSet("Omen", "Spectre", "Envy", "Pavilion")
	case "Acer":
		return randomStringFromSet("Aspire 3", "Aspire 5", "Aspire 7")
	case "Asus":
		return randomStringFromSet("Vivobook", "TUF", "Zenbook")
	case "MSI":
		return randomStringFromSet("Bravo", "Alpha", "GT Titan", "GS Stealth", "GE Raider", "GP Leopard")
	default:
		return ""
	}
}
