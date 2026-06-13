package aws.lambda.base.service;

import com.amazonaws.services.sqs.AmazonSQS;
import com.amazonaws.services.sqs.AmazonSQSClientBuilder;
import com.amazonaws.services.sqs.model.SendMessageRequest;

import lombok.extern.slf4j.Slf4j;

@Slf4j
public class SQS {
    private AmazonSQS sqsClient = AmazonSQSClientBuilder.defaultClient();
    private String queueUrl;

    public void sendMessage(String msg) {
        SendMessageRequest request = new SendMessageRequest().withQueueUrl(queueUrl).withMessageBody(msg);
        sendMessage(request);
    }

    public void sendMessage(String queueUrl, String msg) {
        SendMessageRequest request = new SendMessageRequest().withQueueUrl(queueUrl).withMessageBody(msg);
        sendMessage(request);
    }

    public void sendMessage(SendMessageRequest request) {
        log.info("Send message to {}: {}", request.getQueueUrl(), request.getMessageBody());
        sqsClient.sendMessage(request);

    }
}
