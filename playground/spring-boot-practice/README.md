# Password Validation Service

### Usage
Http request with POST method on `/password/valid` with attribute `passsword` in form.

##### Response format
Check `innova.demo.view.ApiResponse`

### Introduction
In the first version(commit), I use validation api (Default Validation, Validation.validate()) way to implement it
by creating a DTO object and give @Pattern to validate the rules of password.
In this way, the validation can't be replaced/injected by changing configuration.
The last version fully leverage IoC's advantage to implement new validator and replpace/inject
into existing service. Length validation and validation's error message 
can also be injected by giving value in application.properties file.

This service is validation of password. If password is valid,
the controller will return normalized format to the caller without any error message.
Otherwise, it's gonna to throw exception(ValidationFailureException.class).
The exception will be catched by exception handler in controller advice.
The exception handler will response to only specified exception type.
It's a similar way of validation api handle. 

Inside of the validation, since it's multiple validation rule.
I made use of composite design pattern to bundle all smalle rules and
inject every single rules inside of it. When it starts to validate,
it will go through all rules one by one and then give error what it's invalid.

All of validation rules can be controlled in PasswordValidationConfig.class configuration file.
New rules should implement `Validator<String>` and put it under `innova.demo.validator.validation` package.

### Days
##### Analysis and Design
8hr. Check some details and confirm. Design and check some spring document for recall how to code exactly. 

##### Code
8hr, Try and error. Think about how to implement it better.

##### Test
8hr, most of time are used to check annotation and how to make it work between annotations.

##### Document
1hr