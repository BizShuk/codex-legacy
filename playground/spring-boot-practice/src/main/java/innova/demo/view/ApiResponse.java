package innova.demo.view;

import lombok.Builder;
import lombok.Data;
import org.springframework.http.HttpStatus;

import java.util.List;

@Data
@Builder
public class ApiResponse {
    private int status;
    private List errors;
}
