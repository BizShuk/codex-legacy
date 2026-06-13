package innova.demo.service;

import innova.demo.validator.CompositeValidator;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;

import javax.validation.*;

@Service
public class PasswordValidationService {


    @Autowired
    @Qualifier("passwordValidator")
    CompositeValidator<String> passwordValidator;

    public boolean validate(String s) {
        passwordValidator.validate(s);
        return true;
    }

}
