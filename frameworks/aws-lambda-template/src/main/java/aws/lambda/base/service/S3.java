package aws.lambda.base.service;

import java.io.File;
import java.text.MessageFormat;

import com.amazonaws.services.s3.AmazonS3;
import com.amazonaws.services.s3.AmazonS3ClientBuilder;
import com.amazonaws.services.s3.model.GetObjectRequest;
import com.amazonaws.services.s3.model.S3Object;

import aws.lambda.base.util.AWSRegion;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class S3 {
    private final AmazonS3 s3Client = AmazonS3ClientBuilder.standard().withRegion(AWSRegion.getCurrentRegions())
            .build();

    private String bucket;

    public S3(String inputBucket) {
        log.info("Init S3 component with bucket: {}", inputBucket);
    }

    public void putObject(String fileKey, String content) {
        log.info("Upload file to bucket: {}, fileKey: {}", bucket, fileKey);
        s3Client.putObject(bucket, fileKey, content);
    }

    public S3Object getObject(String fileKey) {
        log.info("Get file from bucket:{}, fileKey: {}", bucket, fileKey);
        return s3Client.getObject(bucket, fileKey);
    }

    public void getObject(String fileKey, File dest) {
        log.info("Get S3 object: {}, to {}", fileKey, dest);
        GetObjectRequest request = new GetObjectRequest(bucket, fileKey);
        s3Client.getObject(request, dest);
    }

    @Override
    public String toString() {
        return MessageFormat.format("S3 component with S3 bucket [{0}]", bucket);
    }
}
