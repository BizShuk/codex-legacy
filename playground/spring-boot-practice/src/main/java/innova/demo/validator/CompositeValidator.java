package innova.demo.validator;

import innova.demo.exception.ValidationFailureException;
import innova.demo.validator.validation.Validator;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

@Slf4j
public class CompositeValidator<E> {

    private List<Validator<E>> validators;

    public CompositeValidator (List<Validator<E>> validatorList) {
        this.validators = validatorList;
    }

    public <T> void validate (T t) {
        List<String> errMsgList = new ArrayList<>();

        for (Validator validator: validators) {
            boolean valid = validator.isValid(t);
            if (!valid) {
                errMsgList.add(validator.getErrMsg());
            }
        }

        if (errMsgList.size()>0) throw new ValidationFailureException(errMsgList);
    }
}
