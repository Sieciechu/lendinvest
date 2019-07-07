### Introduction
This app is kind of fast-short, to be done within few days, excercise about playing with golang and DDD and unit testing. It's not meant to be perfect. Everything is done without using a database.

### How to install:
Clone the git repository to folder outside of your $GOPATH as the project uses go modules.

### How to run:
The easiest way is to use docker and prepared scripts:

`$ docker-build-image.sh && docker-run.sh`

Or run

`$ go test && go build -mod vendor app/main.go && ./main`

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

