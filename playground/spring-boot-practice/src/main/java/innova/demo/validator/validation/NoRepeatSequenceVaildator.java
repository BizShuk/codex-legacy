package innova.demo.validator.validation;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

@Component
public class NoRepeatSequenceVaildator implements Validator<String> {
    @Value("${validation.errMsg.noRepeatSequenceVaildator}")
    private String errorMessage;

    @Override
    public boolean isValid(String s) {
        Pattern p = Pattern.compile("^(?!.*(.+)\\1).*$");
        Matcher m = p.matcher(s);
        return m.matches();
    }

    @Override
    public String getErrMsg() {
        return errorMessage;
    }
}
