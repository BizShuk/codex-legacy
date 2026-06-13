# Practice Project

### Requirement:
⦁ Write a password validation service, meant to be configurable via IoC (using dependency injection engine of your choice)
⦁ Must consist of a mixture of lowercase letters and numerical digits only, with at least one of each.
⦁ Must be between 5 and 12 characters in length.
⦁ Must not contain any sequence of characters immediately followed by the same sequence.


### Question: 
1. IoC, any configurable wya should be fine, xml or configuration class file?
2. Should be a class with a String input only? No need for a whole system
3. How to put input param with type of validator? IoC to inject a validator inside PasswordValidateService
4. What is return value for this service?
5. Does it require to return all failed reason, not just return one failed reason once failed?


### Design Thought:
- each rule should be a class
- all rules have the same interface
- a validator will be a composition of bunch rules
- different validator has the same interface
- the service will call validator interface


### Project Details:
- use only spring context instead of all spring framework
- Main class load IoC and try simple one
- Include test in this case
- make test extentable with different rule and validators
- 
