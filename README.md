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
go run main.go input.sample.JSON.yaml
```
