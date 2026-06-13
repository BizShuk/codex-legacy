package innova.demo.controller;

import innova.demo.service.PasswordValidationService;
import innova.demo.service.PasswordValidationService;
import innova.demo.view.ApiResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
public class PasswordValidationController {

    @Autowired
    private PasswordValidationService passwordValidationService;

    @PostMapping("/password/valid")
    @ResponseStatus(HttpStatus.OK)
    @ResponseBody
    public ApiResponse validPassword (@RequestParam("password") String password) {
        passwordValidationService.validate(password);
        return ApiResponse.builder().status(HttpStatus.OK.value()).build();
    }
}
