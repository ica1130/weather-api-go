# Weather API in Go

**Weather API** practice project, a RESTful API built in Go for fetching, caching, and serving weather data! This API utilizes the [Visual Crossing Weather API](https://www.visualcrossing.com/weather-api) for real-time weather data and employs Redis for efficient caching. Additionally, a rate limiter is implemented to manage API request traffic and ensure optimal performance.

![image](https://github.com/user-attachments/assets/c4a7a1eb-e25c-4ca9-90c2-caea5d5a7077)

## Features

- **Weather Data**: Retrieves detailed weather information including temperature, humidity, wind speed, and more.
- **Caching with Redis**: Redis is used to cache weather data, reducing the number of API calls to Visual Crossing and enhancing response speed.
- **Rate Limiting**: A rate limiter is in place to control the volume of requests per user, protecting the API from excessive usage and ensuring stability.

## Project Specifics

- Go (1.18 or later)
- Redis server
- Visual Crossing API Key (sign up for a key [here](https://www.visualcrossing.com/weather-api))

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ica1130/weather-api-go.git
   cd weather-api-go
   ```

2. Install Go dependencies:
   ```bash
   go mod download
   ```

3. Set up your environment variables:

   Create a `.env` file in the root directory with the following keys:
   ```plaintext
   WEATHER_API_KEY=<your_visual_crossing_api_key>
   ```

4. Run the API server:
   ```bash
   go run ./cmd/api
   ```

## Endpoints

### Get Weather Data

Fetch weather information by providing the location.

```http
GET http://localhost:4000/v1/weather?city={city}
```

#### Parameters

- `city` (string) - The name of the city or location to fetch weather data for.

## API Key Configuration

The **Visual Crossing API Key** is required for fetching weather data. Configure your key in the `.env` file to avoid hardcoding it.

## License

This project is licensed under the MIT License.
