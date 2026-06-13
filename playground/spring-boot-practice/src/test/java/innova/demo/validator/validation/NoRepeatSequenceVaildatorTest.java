package innova.demo.validator.validation;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.TestPropertySource;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import static org.junit.jupiter.api.Assertions.*;

@ExtendWith(SpringExtension.class)
@TestPropertySource("/application.properties")
@ContextConfiguration(classes = NoRepeatSequenceVaildator.class)
class NoRepeatSequenceVaildatorTest {
    @Value("${validation.errMsg.noRepeatSequenceVaildator}")
    private String errorMsg;

    @Autowired
    NoRepeatSequenceVaildator nValidator;

    @Test
    void Should_Pass_When_NoRepeatSequence() {
        assertFalse(nValidator.isValid("abcd1abcd1"));
    }

    @Test
    void Should_Fail_When_RepeatSequenceExist() {
        assertFalse(nValidator.isValid("abcd1abcd1"));
    }

    @Test
    void Should_ReturnErrorMessage() {
        assertEquals(errorMsg, nValidator.getErrMsg());
    }
}