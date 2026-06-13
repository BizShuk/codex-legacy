package aws.lambda.base.service;

import java.util.HashMap;
import java.util.Map;

import com.amazonaws.services.secretsmanager.AWSSecretsManager;
import com.amazonaws.services.secretsmanager.AWSSecretsManagerClientBuilder;
import com.amazonaws.services.secretsmanager.model.GetSecretValueRequest;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;

import aws.lambda.base.util.AWSRegion;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class SecretManager {
    private AWSSecretsManager secretsManagerClient = AWSSecretsManagerClientBuilder.standard()
            .withRegion(AWSRegion.getCurrentRegions()).build();

    private TypeReference<Map<String, Object>> typeRef = new TypeReference<Map<String, Object>>() {
    };

    public Map<String, Object> fetchSecrets(String secretId) {
        log.debug("fetchSecrets: {}", secretId);
        Map<String, Object> resultInMap = new HashMap<>();

        try {
            GetSecretValueRequest request = new GetSecretValueRequest().withSecretId(secretId);

            String secrets = secretsManagerClient.getSecretValue(request).getSecretString();
            ObjectMapper mapper = new ObjectMapper();
            resultInMap = mapper.readValue(secrets, typeRef);
        } catch (Exception e) {
            // Somewhere might have it instead of fetching from SecretsManager
            log.error("No such secrets in Secrets Manager: {}", secretId);
        }
        return resultInMap;
    }
}
