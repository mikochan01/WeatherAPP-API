// Ambil data cuaca dari API
const apiKey = '63c031c28f3e4ce3b0093142230811';
const city = 'Tangerang';
const apiUrl = `http://api.weatherapi.com/v1/forecast.json?key=63c031c28f3e4ce3b0093142230811&q=Tangerang&days=1&aqi=no&alerts=no`;

// Fungsi untuk mengambil data dari API
async function fetchData() {
    try {
        const response = await fetch(apiUrl);
        const data = await response.json();
        displayWeather(data);
    } catch (error) {
        console.error('Error fetching data:', error);
    }
}

// Fungsi untuk menampilkan informasi cuaca ke antarmuka pengguna
function displayWeather(data) {
    const weatherContainer = document.getElementById('weather-container');
    const weatherCard = document.createElement('div');
    weatherCard.className = 'weather-card';

    const temperature = data.current.temp_c;
    const humidity = data.current.humidity;
    const condition = data.current.condition.text;

    weatherCard.innerHTML = `
        <h2>Weather in ${city}</h2>
        <p>Temperature: ${temperature}°C</p>
        <p>Humidity: ${humidity}%</p>
        <p>Condition: ${condition}</p>
    `;

    weatherContainer.appendChild(weatherCard);
}

// Panggil fungsi fetchData untuk mendapatkan data dan menampilkan informasi cuaca
fetchData();
