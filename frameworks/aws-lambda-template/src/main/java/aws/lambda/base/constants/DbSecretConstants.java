package aws.lambda.base.constants;

public class DbSecretConstants {
    public static final String API_SECRET_SUFFIX = "/api";

    private DbSecretConstants() {
        throw ExceptionConstants.ILLEGAL_STATE_EXCEPTION;
    }
}
