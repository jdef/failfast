## usage

```sh
:; ./demo.sh

go version
go version go1.10 linux/amd64
go test -c -o failfast_testbin
./failfast_testbin -test.v -test.failfast
=== RUN   TestA
--- FAIL: TestA (5.00s)
        failfast_test.go:10: failed test
=== RUN   TestB
--- PASS: TestB (1.00s)
FAIL
```
