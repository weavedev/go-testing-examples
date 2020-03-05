# Golang testing examples

This repositories illustrates some examples of creating unit tests in Golang using the Testify library.

All tests can be run by executing `make test`. 

When the tests ran, function coverage can be shown by executing `make funcCoverage`

When the tests ran, you can open the coverage report by executing `make coverageReport`

## cmd

The cmd folder contains three examples for using tests.
1. The fibonacci-test-example shows some simple unit test example and a benchmark example
2. The parallel-test-example shows how to run testify suites in parallel
3. The api-test-example demonstrates some examples on how to use mocks in tests

## packages

The packages contain packages used in the test examples with the exception of the models package, which illustrates how to execute tests on top of a Postgres database.