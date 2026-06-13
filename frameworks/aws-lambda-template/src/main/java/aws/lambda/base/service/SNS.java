package aws.lambda.base.service;

import com.amazonaws.services.sns.AmazonSNS;
import com.amazonaws.services.sns.AmazonSNSClientBuilder;
import com.amazonaws.services.sns.model.PublishRequest;
import com.amazonaws.services.sns.model.PublishResult;

import lombok.extern.slf4j.Slf4j;

@Slf4j
public class SNS {
    private final AmazonSNS snsClient = AmazonSNSClientBuilder.standard().build();

    public PublishResult publish(String topic, String msg) {
        log.info("Send to SNS topic: {}, message: {}", topic, msg);
        final PublishRequest request = new PublishRequest(topic, msg);
        final PublishResult result = snsClient.publish(request);
        log.info("Received SNS response: {}", result);
        return result;
    }
}
