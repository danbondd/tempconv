package converter_test

import (
	"fmt"
	"testing"

	"github.com/danbondd/temperature/converter"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	fahrenheit := converter.CelsiusToFahrenheit(20)

	if fahrenheit != 68 {
		t.Errorf("expected 68, got: %f", fahrenheit)
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	kelvin := converter.CelsiusToKelvin(30)

	if kelvin != 303.2 {
		t.Errorf("expected 303.2, got: %f", kelvin)
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	celsius := converter.FahrenheitToCelsius(55)

	if celsius != 12.8 {
		t.Errorf("expected 12.8, got: %f", celsius)
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	kelvin := converter.FahrenheitToKelvin(50)

	if kelvin != 283.2 {
		t.Errorf("expected 283.2, got: %f", kelvin)
	}
}

func TestKelvinToCelsius(t *testing.T) {
	celsius := converter.KelvinToCelcius(300)

	if celsius != 26.9 {
		t.Errorf("expected 26.9, got: %f", celsius)
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	fahrenheit := converter.KelvinToFahrenheit(300)

	if fahrenheit != 80.3 {
		t.Errorf("expected 80.3, got: %f", fahrenheit)
	}
}

func TestCelsiusToString(t *testing.T) {
	celsius := fmt.Sprintf("%s", converter.Celsius(12))

	if celsius != "12°C" {
		t.Errorf("expected 12°C, got: %s", celsius)
	}
}

func TestFahrenheitToString(t *testing.T) {
	fahrenheit := fmt.Sprintf("%s", converter.Fahrenheit(55))

	if fahrenheit != "55°F" {
		t.Errorf("expected 55°F, got: %s", fahrenheit)
	}
}

func TestKelvinToString(t *testing.T) {
	kelvin := fmt.Sprintf("%s", converter.Kelvin(300))

	if kelvin != "300K" {
		t.Errorf("expected 300K, got: %s", kelvin)
	}
}
