package innova.demo.validator.validation;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

@Component
public class LowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator implements Validator<String> {
    @Value("${validation.errMsg.lowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator}")
    private String errorMessage;

    public boolean isValid(String s) {
        Pattern p = Pattern.compile("^(?=.*\\d)(?=.*[a-z])[a-z0-9]+$");
        Matcher m = p.matcher(s);
        return m.matches();
    }

    @Override
    public String getErrMsg() {
        return errorMessage;
    }

}
