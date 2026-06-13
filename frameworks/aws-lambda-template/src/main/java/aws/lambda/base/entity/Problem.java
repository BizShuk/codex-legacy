package aws.lambda.base.entity;

import lombok.Data;

@Data
public class Problem {
    private String title;
    private String status;
    private String detail;
}
