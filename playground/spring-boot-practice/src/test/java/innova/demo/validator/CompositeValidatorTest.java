package innova.demo.validator;

import innova.demo.config.PasswordValidationConfig;
import innova.demo.exception.ValidationFailureException;
import innova.demo.validator.validation.LowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator;
import innova.demo.validator.validation.NoRepeatSequenceVaildator;
import innova.demo.validator.validation.PasswordLengthValidator;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.when;

@ExtendWith(SpringExtension.class)
@ContextConfiguration(classes = PasswordValidationConfig.class)
class CompositeValidatorTest {
    @MockBean
    LowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator lValidator;
    @MockBean
    NoRepeatSequenceVaildator nValidator;
    @MockBean
    PasswordLengthValidator pValidator;

    @Test
    void Should_Pass() {
        when(lValidator.isValid(anyString())).thenReturn(true);
        when(nValidator.isValid(anyString())).thenReturn(true);
        when(pValidator.isValid(anyString())).thenReturn(true);
        CompositeValidator compositeValidator = new CompositeValidator<String>(
                Arrays.asList(
                        lValidator,
                        nValidator,
                        pValidator
                )
        );

        compositeValidator.validate(anyString());
    }

    @Test
    void Should_ThrowException() {
        when(lValidator.isValid(anyString())).thenReturn(false);
        when(nValidator.isValid(anyString())).thenReturn(false);
        when(pValidator.isValid(anyString())).thenReturn(false);
        CompositeValidator compositeValidator = new CompositeValidator<String>(
                Arrays.asList(
                        lValidator,
                        nValidator,
                        pValidator
                )
        );

        ValidationFailureException ex = assertThrows(
                ValidationFailureException.class,
                () -> compositeValidator.validate(anyString())
        );
        assertEquals(3, ex.getErrMsgList().size());
    }

}