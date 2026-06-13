package innova.demo.validator.validation;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.TestPropertySource;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import java.text.MessageFormat;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;

@ExtendWith(SpringExtension.class)
@TestPropertySource("/application.properties")
@ContextConfiguration(classes = PasswordLengthValidator.class)
class PasswordLengthValidatorTest {
    @Value("${validation.errMsg.passwordLengthValidator}")
    private String errorMsg;

    @Autowired
    PasswordLengthValidator pValidator;

    @Value("${password.length.min:5}")
    private int minLength;

    @Value("${password.length.max:12}")
    private int maxLength;

    @Test
    void Should_Pass_When_LengthInRange() {
        assertFalse(pValidator.isValid("abc1"));
    }

    @Test
    void Should_Fail_When_LengthLessThan5() {
        assertFalse(pValidator.isValid("abc1"));
    }

    @Test
    void Should_Fail_When_LengthMoreThan12() {
        assertFalse(pValidator.isValid("abcd1819njedjhaf"));
    }

    @Test
    void Should_ReturnErrorMessage() {
        assertEquals(MessageFormat.format(errorMsg, minLength, maxLength), pValidator.getErrMsg());
    }
}