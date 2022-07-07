# UnitTest

## Name of file
* Ending a file's name with _test.go tells the go test command that this file contains test functions.

## commands
* testing.T (Unit Test)
  * t.Fail(), will fail unit test, but keep execute unit test
  * t.FailNow(), will fail unit test immediately, without continuing execution
  * t.Error(args…), working like t.Fail() with adding log errors
  * t.Fatal(args…), working like t.FailNow() with adding log errors 
* testing.M (Life-cycle Testing)
* testing.B (Benchmarking)
