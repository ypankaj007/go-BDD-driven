# Behaviour-Driven Development in GO

### What is BDD?
BDD is a way for software teams to work that closes the gap between business people and technical people by:

* Encouraging collaboration across roles to build shared understanding of the problem to be solved
* Working in rapid, small iterations to increase feedback and the flow of value
* Producing system documentation that is automatically checked against the system’s behaviour

We do this by focusing collaborative work around concrete, real-world examples that illustrate how we want the system to behave. We use those examples to guide us from concept through to implementation, in a process of continuous collaboration.
[reference ](https://cucumber.io/)

### [Repo example](https://github.com/cucumber/godog):
In the example, we have an api which returns user list created by using go http package.
For this simple example, we have api.feature file inside feature folder which contains scenario outlines steps.

### Run the test case
```sh
$ godog features/user.feature
```
