package aws.lambda.base.exception;

public class BusinessLogicException extends RuntimeException {
    private static final long serialVersionUID = 4735162649209980282L;

    public BusinessLogicException(String msg) {
        super(msg);
    }
}
