package innova.demo.config;

import innova.demo.validator.validation.LowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator;
import innova.demo.validator.validation.NoRepeatSequenceVaildator;
import innova.demo.validator.validation.PasswordLengthValidator;
import innova.demo.validator.CompositeValidator;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.Arrays;

@Configuration
public class PasswordValidationConfig {
    @Autowired
    private LowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator lowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidation;
    @Autowired
    private NoRepeatSequenceVaildator noSameSequenceVaildator;
    @Autowired
    private PasswordLengthValidator passwordLengthValidator;

    @Bean("passwordValidator")
    public CompositeValidator getPasswordValidator () {
        return new CompositeValidator<String>(Arrays.asList(
                lowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidation,
                noSameSequenceVaildator,
                passwordLengthValidator
        ));
    }
}
