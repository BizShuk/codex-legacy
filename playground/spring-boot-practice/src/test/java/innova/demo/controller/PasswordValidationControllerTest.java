package innova.demo.controller;

import innova.demo.exception.ValidationFailureException;
import innova.demo.service.PasswordValidationService;
import lombok.extern.slf4j.Slf4j;
import org.hamcrest.core.IsNull;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.junit.jupiter.SpringExtension;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;
import org.springframework.test.web.servlet.RequestBuilder;

import java.util.Arrays;

import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.BDDMockito.given;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@Slf4j
@ExtendWith(SpringExtension.class)
@WebMvcTest(PasswordValidationController.class)
class PasswordValidationControllerTest {

    @Autowired
    private MockMvc mvc;

    @MockBean
    private PasswordValidationService passwordValidationService;

    @Test
    void Should_ReturnStatusOk_When_ValidationPassed() throws Exception {

        RequestBuilder request = post("/password/valid")
                .param("password", "abcd1");

        mvc.perform(request)
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.status").value("200"))
                .andExpect(jsonPath("$.errors").value(IsNull.nullValue()));
    }

    @Test
    void Should_ThrowException_When_ValidationNotPassed() throws Exception {
        given(passwordValidationService.validate(anyString())).willThrow(
                new ValidationFailureException()
        );

        RequestBuilder request = post("/password/valid")
                .param("password", anyString());

        mvc.perform(request)
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("$.status").value("400"));
    }
}