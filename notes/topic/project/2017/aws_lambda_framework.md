# Lambda framework

Java is an awesome language. It fills with comprehensive framework and libraries before cloud solution showing up. Faster growing cloud solutions give more easier and concise ways to do exactly same thing in old time. Aws Lambda is one of them. Just give a function with business logic, it will perform as a service without configure of MVC or HTTP server. In the beginning, people were just moving all existing staff on the cloud. But it is not working like that. Take AWS Lambda as example, MVC framework just makes the code base worse. So, where should we put those non-functional requirement, such as authentication, secure parameters? The Lambda framework is coming out my mind, just like SprintBoot. It has it own mechanism and flow to process steps by steps, loading properties from AWS Secret Manager, initializing AWS components, connecting database. It leverages abstraction class and extending from this class. Every things will work behind the scene by just configuring properties. And you just need to take care of business logic as AWS Lambda design to do in orignal.

### Dependancy Reduction

Java dependancies are heavy. It takes time to remove unnecessary from bundled package.

### DB Connection

Same as dependancy issue. Using single independant library instead of high level DB library is hard to make it scale/flexible. [Data Souce] [DB Driver]
