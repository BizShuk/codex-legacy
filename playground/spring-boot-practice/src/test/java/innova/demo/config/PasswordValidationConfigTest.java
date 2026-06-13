package innova.demo.config;

import innova.demo.validator.CompositeValidator;
import innova.demo.validator.validation.LowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator;
import innova.demo.validator.validation.NoRepeatSequenceVaildator;
import innova.demo.validator.validation.PasswordLengthValidator;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import static org.junit.jupiter.api.Assertions.*;

@ExtendWith(SpringExtension.class)
@ContextConfiguration(classes = {
        PasswordValidationConfig.class,
        LowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator.class,
        NoRepeatSequenceVaildator.class,
        PasswordLengthValidator.class
})
class PasswordValidationConfigTest {
    @Autowired
    LowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator lValidator;
    @Autowired
    NoRepeatSequenceVaildator nValidator;
    @Autowired
    PasswordLengthValidator pValidator;

    @Autowired
    @Qualifier("passwordValidator")
    CompositeValidator passwordValidator;

    @Test
    void Should_BeNotNull_When_LoadFromConfigFile() {
        assertNotNull(lValidator);
        assertNotNull(nValidator);
        assertNotNull(pValidator);
        assertNotNull(passwordValidator);
    }
}