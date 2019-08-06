# test-case-engine
An engine to create test data of any kind!

The application takes a filename as an argument, which should point to a YAML file with the following structure:

* __title__: A name for the data that you are generating. If you are generating SQL, then this will also be used as the table name in your insert statements.
* __output__: The format that will be used to generate your test data.  Currently only __JSON__ and __SQL__ are supported
* __total__: The total amount of items that will be generated
* __fields__: Definitions for the different pieces of data that will be generated

The __fields__ section can have one or more items with the following structure

* __name__: The name of the data that is being generated
* __type__: The type of data being generated.  Currently only __alpha__ and __integer__ are supported
* __min__: For alpha fields, this is the minimum length of the string that's generated.  For integer fields this is the smallest number that can be generated
* __max__: For alpha fields, this is the maximum length of the string that's generated.  For integer fields this is the largest number that can be generated

There is a sample input file included in the repository.  See input.sample.yaml for details.

To run the application with the sample input file, run the following command from the `/src` directory:

```sh
go run main.go input.sample.yaml
```

## SonarQube and Testing
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

### Test case scanning
To run unit tests, run the following command:
```sh
go test testing/
```
