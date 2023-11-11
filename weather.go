package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetWeather() (string, error) {
	// Ganti dengan kunci API dan kota yang sesuai
	apiKey := "63c031c28f3e4ce3b0093142230811"
	city := "Tangerang"

	// Membuat URL dengan parameter kota
	apiURL := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", apiKey, city)

	// Membuat permintaan HTTP GET
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Mengecek apakah respons berhasil (status kode 200 OK)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error: %s", resp.Status)
	}

	// Mendekode respons JSON
	var weatherData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		return "", fmt.Errorf("Error decoding JSON: %v", err)
	}

	// Mendapatkan informasi suhu, kelembapan, dan kondisi cuaca
	temperature := weatherData["current"].(map[string]interface{})["temp_c"].(float64)
	humidity := weatherData["current"].(map[string]interface{})["humidity"].(float64)
	condition := weatherData["current"].(map[string]interface{})["condition"].(map[string]interface{})["text"].(string)

	// Menampilkan informasi cuaca
	result := fmt.Sprintf("Weather in %s:\n", city)
	result += fmt.Sprintf("Temperature: %.2f°C\n", temperature)
	result += fmt.Sprintf("Humidity: %.2f%%\n", humidity)
	result += fmt.Sprintf("Condition: %s\n", condition)

	return result, nil
}
