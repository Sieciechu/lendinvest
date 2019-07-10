### Introduction
This is just an *excercise*. The library is kind of fast-short, to be done within few days, excercise about playing with golang and DDD and unit testing. It's not meant to be perfect. Everything is done without using a database. It's more a package than app. App is non interactive and is created only to show basics of package. The most things happen in the tests.

### How to install:
Clone the git repository to folder outside of your $GOPATH as the project uses go modules.

### How to run:
The easiest way is to use docker and prepared scripts:

`$ docker-build-image.sh && docker-run.sh`

Or run

`$ go test && go build -mod vendor app/main.go && ./main`

### Link to GoDoc
[https://godoc.org/github.com/Sieciechu/lendinvest](https://godoc.org/github.com/Sieciechu/lendinvest)

### More details about the exercise
Exercise

The app is about connecting people who want to invest their money, with investments to
those who want to borrow. One of the parts of a business is to give investors a way to invest in a loan for
them to earn a return (monthly interest payment).

Model
- Each of our loans has a start date and an end date.
- Each loan is split in multiple tranches.
- Each tranche has a different monthly interest percentage.
- Also each tranche has a maximum amount available to invest. So once the maximum is reached, further investments can't be made in that tranche.
- As an investor, I can invest in a tranche at any time if the loan it’s still open, the maximum
available amount was not reached and I have enough money in my virtual wallet.
- At the end of the month we need to calculate the interest each investor is due to be paid.

Scenario
- Given a loan (start 2019-10-01 end 2019-11-15).
- Given the loan has 2 tranches called A and B (3% and 6% monthly interest rate) each with
1,000 pounds amount available.
- Given each investor has 1,000 pounds in his virtual wallet.
- As “Investor 1” I’d like to invest 1,000 pounds on the tranche “A” on 03/10/2019: “ok”.
- As “Investor 2” I’d like to invest 1 pound on the tranche “A” on 04/10/2019: “exception”.
- As “Investor 3” I’d like to invest 500 pounds on the tranche “B” on 10/10/2019: “ok”.
- As “Investor 4” I’d like to invest 1,100 pounds on the tranche “B” 25/10/2019: “exception”.
- On 01/11/2019 the system runs the interest calculation for the period 01/10/2019 ->
31/10/2019:
- “Investor 1” earns 28.06 pounds
- “Investor 3” earns 21.29 pounds

### Some private conclusions, thoughts from this excercise:
* `$ go mod init github.com/Sieciechu/lendinvest` allowed to work outside of $GOPATH - see `$ go help modules`
* Though I have no other dependencies `$ go mod vendor; go build -mod vendor` allows to keep dependencies in vendor folder with needed versions
* I come from PHP environment. Golang basic syntax is easy and there was no trouble to model domain for **this topic** in golang.
* Testing most of functions was quite handy. Lack of object mocking in golang testing package dissallowed me to test some functionalities at the top level. I would have to search for some testing framework (if go supports somehow mocking objects by reflection) or would have to introduce some more interfaces so I could implement mock/dummy methods.
* As go's testing package allows to test non-exported/support functions/methods I got excited and tested them too, and forgot about discipline to focus more on kind of public ones/main functionalities. Which this time lead to: though I have same amount of things tested, but some of them are not so well structured - few test cases of single functionality are covered in "sub-methods"/"support-methods".
* I wanted to create other array having pointers to the original array values. So I wrote something like:
```go
// ...
otherArr := make([]*SomeStruct, 0)
for _, p := range arr {
	otherArr = append(otherArr, &p)
}
// ...
```
 resulted all elements of otherArr pointed to last p. Later thought about it: p is a **new single variable** and contains **copy** of array value, add I passed address of p, not address of arr's element. The fix was to use classical for loop with index, to pass "pointers to arr values".
* Most of the times I used slices as containers, just in the end I found go's container list (https://golang.org/pkg/container/), maybe it would be more handy/better to use
* Feared of the lack of exceptions. But when you keep in mind, that exceptions from lower levels should not leak up just as they are to upper levels, so anyway you have to process cought lower level exceptions. For example: the end user of a game should see error like 'error in saving game' rather than 'sql exception, state 20090 in file /path/to/file line 30'. So in this exercise handling errors each time instead of catching exceptions was not so annoying as I had expected (at least in this exercise).
* Accidental assignment to nil map caused **runtime** panic. Simple example: `var people map[string]int;    people["john"] = 32`. So remember to initialize it with make or with literals
* It looks that structs: Loan & Tranche could implement same interface for MakeInvestment, so in future it would be possible to invest in something else implementing this interface
* Visual Studio Code is ok, but not so good as JetBrains Goland ;).
