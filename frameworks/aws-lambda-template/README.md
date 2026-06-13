# AwsLambdaBaseTemplate

The mechanism is that the lambda handler needs to extend template class and the template class will initializa LambdaConfig and LambdaContext. LambdaConfig and LambdaContext are responsible for SecretManager (loading property) and other selected components. And, in bindComponent method, you can bind those components into the handler class by `ctx.getProperty(LambdaConfig.Components.*)`

### Hot To Import
```
<dependency>
    <groupId>com.chc.common</groupId>
    <artifactId>AwsLambdaBaseTemplate</artifactId>
    <version>{version}</version>
</dependency>
```

### Simple Tutorial
Extend `AbstractLambdaTemplate.java` and override necessary method. The main functionality should be in `handlerService` method.

### DB
DB scheme defined by `getProductName` in Lambda function which extend `AbstractLambdaTempalte.java`


### Secret list from Secret Manager
`getScretList` is one of method necessary to be overrode. It just simply need to return the secret string list. With `${var}` inside, the variable will be replaced with env variables. If profile variable from lambda environment variable is dev, `"lambda/secret/${profile}"` => `"lambda/secret/dev"`

### Selected Components
- S3
- SNS
- SQS

### Local Testing
Once it is ready, you can run it with given enviroment variable. Otherwise, the secret list won't replace with placeholder. Of course, you have to define your main function and give arguments for it.

### HttpClient
Recommend `Uni-Rest Java` Check [Here](http://kong.github.io/unirest-java)


### TODO
- Cache secrets from Secret Manager
- reconnect db with refresh db scret when some connection problem happens