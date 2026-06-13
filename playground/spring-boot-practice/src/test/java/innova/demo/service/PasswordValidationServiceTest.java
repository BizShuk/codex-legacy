package innova.demo.service;


import innova.demo.exception.ValidationFailureException;
import innova.demo.validator.CompositeValidator;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.doNothing;
import static org.mockito.Mockito.doThrow;

@ExtendWith(SpringExtension.class)
@ContextConfiguration(classes = PasswordValidationService.class)
class PasswordValidationServiceTest {

    @MockBean
    @Qualifier("passwordValidator")
    CompositeValidator<String> passwordValidator;

    @Autowired
    PasswordValidationService passwordValidationService;

    @Test
    void Should_Pass_When_ValidationPassed() {
        doNothing().when(passwordValidator).validate(anyString());
        boolean bool = passwordValidationService.validate(anyString());
        assertTrue(bool);
    }

    @Test
    void Should_ThrowException_When_ValidationNotPassed() {
        doThrow(new ValidationFailureException()).when(passwordValidator).validate(anyString());
        assertThrows(ValidationFailureException.class, () -> {
            boolean bool = passwordValidationService.validate(anyString());
        });
    }
}