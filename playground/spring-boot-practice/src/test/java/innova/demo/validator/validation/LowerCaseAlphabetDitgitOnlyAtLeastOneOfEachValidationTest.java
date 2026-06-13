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
@ContextConfiguration(classes = LowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator.class)
class LowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidationTest {
    @Value("${validation.errMsg.lowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator}")
    private String errorMsg;

    @Autowired
    LowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator lValidator;

    @Test
    void Should_Pass_When_InputValidated() {
        assertTrue(lValidator.isValid("abc123"));
    }

    @Test
    void Should_Fail_When_AnyLowercaseNotExist() {
        assertFalse(lValidator.isValid("123123789"));
    }

    @Test
    void Should_Fail_When_AnyDigitNotExist() {
        assertFalse(lValidator.isValid("abcdef"));
    }

    @Test
    void Should_Fail_When_NotLowercaseDigitExist() {
        assertFalse(lValidator.isValid("ASDJIOMW!@#*()*"));
    }

    @Test
    void Should_ReturnErrorMessage() {
        assertEquals(errorMsg, lValidator.getErrMsg());
    }
}