package innova.demo;

import innova.demo.view.ApiResponse;
import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;

import java.text.MessageFormat;
import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

@Slf4j
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class DemoApplicationTests {
	@LocalServerPort
	private int port;

	@Autowired
	private TestRestTemplate restTemplate;

	@Value("${validation.errMsg.lowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidator}")
	private String lowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidatorErrorMsg;
	@Value("${validation.errMsg.noRepeatSequenceVaildator}")
	private String noRepeatSequenceVaildatorErrorMsg;
	@Value("${validation.errMsg.passwordLengthValidator}")
	private String passwordLengthValidatorErrorMsg;
	@Value("${password.length.min}")
	private int passwordMinLength;
	@Value("${password.length.max}")
	private int passwordMiaxLength;


	@Test
	void Should_ReturnOkStatus_When_ValidationPassed() {

		HttpHeaders headers = new HttpHeaders();
		headers.setContentType(MediaType.APPLICATION_FORM_URLENCODED);

		MultiValueMap<String, String> map= new LinkedMultiValueMap<String, String>();
		map.add("password", "abcde123");

		HttpEntity<MultiValueMap<String, String>> request = new HttpEntity<>(map, headers);

		ApiResponse response = this.restTemplate.postForObject(
				"http://localhost:" + port + "/password/valid",
				request,
				ApiResponse.class
				);
		log.info("result: {}", response);
		assertEquals(200, response.getStatus());
		assertEquals(null, response.getErrors());
	}

	@Test
	void Should_ReturnBadRequestStatusWithErrors_When_ValidationFailed() {

		HttpHeaders headers = new HttpHeaders();
		headers.setContentType(MediaType.APPLICATION_FORM_URLENCODED);

		MultiValueMap<String, String> map= new LinkedMultiValueMap<String, String>();
		map.add("password", "ABAB");

		HttpEntity<MultiValueMap<String, String>> request = new HttpEntity<>(map, headers);

		ApiResponse response = this.restTemplate.postForObject(
				"http://localhost:" + port + "/password/valid",
				request,
				ApiResponse.class
		);
		log.info("result: {}", response);
		assertEquals(400, response.getStatus());
		assertEquals(3, response.getErrors().size());

		List<String> errorList = response.getErrors();
		assertTrue(errorList.contains(lowerCaseAlphabetDitgitOnlyAtLeastOneOfEachValidatorErrorMsg));
		assertTrue(errorList.contains(noRepeatSequenceVaildatorErrorMsg));
		assertTrue(errorList.contains(
				MessageFormat.format(passwordLengthValidatorErrorMsg,passwordMinLength,passwordMiaxLength)
		));
	}
}
