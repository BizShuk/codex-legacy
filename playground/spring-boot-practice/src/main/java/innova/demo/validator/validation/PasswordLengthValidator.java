package innova.demo.validator.validation;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.text.MessageFormat;

@Component
public class PasswordLengthValidator implements Validator<String> {
    @Value("${validation.errMsg.passwordLengthValidator}")
    private String errorMessage;

    @Value("${password.length.min:5}")
    private int minLength;

    @Value("${password.length.max:12}")
    private int maxLength;

    public boolean isValid(String s) {
        return s.length() >= minLength && s.length() <= maxLength;
    }

    @Override
    public String getErrMsg() {
        return MessageFormat.format(errorMessage, minLength, maxLength);
    }
}
