# SonarQube and Testing
For convenience a `docker-compose.yaml` file is included in this repository which will standup a SonarQube and PostgreSQL container to use in code scanning.  To use this service, you will need to set the following environment variables:
```sh
export SONAR_POSTGRES_USER=<your username>
export SONAR_POSTGRES_PASSWORD=<your password>
```

Once that's set you can run the service:
```sh
docker-compose up -d
```

SonarQube will be running at http://localhost:9000

You will need `sonar-scanner` running locally. To do that follow the directions here: https://docs.sonarqube.org/latest/analysis/scan/sonarscanner/

Once your local SonarQube instance is running and `sonar-scanner` is installed you can analyze this project from the project root by simply running:
```sh
sonar-scanner
```

## Test case scanning
To run unit tests, run the following command:
```sh
go test testing/
```
