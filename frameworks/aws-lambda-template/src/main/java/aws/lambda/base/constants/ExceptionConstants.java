package aws.lambda.base.constants;

public final class ExceptionConstants {

    private ExceptionConstants() {
        throw ILLEGAL_STATE_EXCEPTION;
    }

    public static final IllegalStateException ILLEGAL_STATE_EXCEPTION = new IllegalStateException("Utility class");
}
