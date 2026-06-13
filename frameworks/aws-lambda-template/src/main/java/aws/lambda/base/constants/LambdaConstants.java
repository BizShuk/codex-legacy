package aws.lambda.base.constants;

public final class LambdaConstants {
    private LambdaConstants() {
        throw ExceptionConstants.ILLEGAL_STATE_EXCEPTION;
    }

    public static final String APPNAME_PROPERTY = "lambda.app.name";
    public static final String DB_FORCESSL_PROPERTY = "lambda.datasource.forceSSL";
    public static final String TS_PASSWORD_PROPERTY = "app.cert.ts.password";
    public static final String LAMBDA_TS_LOCAL_PATH = "/tmp/ssl/cert/app.ts";
}
