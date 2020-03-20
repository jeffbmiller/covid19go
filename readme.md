# COVID-29 REST API Example

A RESTful API example for getting Covid-19 statistics all around written in GO.

## Installaton & Run
1. Clone Repository 
```
git clone https://github.com/jeffbmiller/covid19go.git
```
2. Inside the directory run this command to build the docker image
```
docker build -t covid19go .
```
3. Run the Docker image
```
docker run -d -p 5000:8080 covid19go
```
4. Open your web browser and go to http://localhost:5000/countries
 http://localhost:5000/countries/<countryname> will show stats for that country only