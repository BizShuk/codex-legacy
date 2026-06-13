package innova.demo.controller.advice;

import innova.demo.exception.ValidationFailureException;
import innova.demo.view.ApiResponse;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import javax.validation.ConstraintViolationException;

import java.util.stream.Collectors;

@ControllerAdvice
@ResponseBody
public class ControllerExceptionHandler extends ResponseEntityExceptionHandler {

    @ExceptionHandler(ValidationFailureException.class)
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    public ApiResponse handleValidationFailureException (
            ValidationFailureException ex,
            WebRequest request
    ) {
        ApiResponse apiError = ApiResponse.builder()
                .status(HttpStatus.BAD_REQUEST.value())
                .errors(ex.getErrMsgList())
                .build();

        return apiError;
    }

}
