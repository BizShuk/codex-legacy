package aws.lambda.base.config;

import java.io.File;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Map;

import com.amazonaws.services.lambda.runtime.Context;
import com.mysql.cj.util.StringUtils;

import org.apache.commons.text.StringSubstitutor;

import aws.lambda.base.constants.DbSecretConstants;
import aws.lambda.base.constants.LambdaConstants;
import aws.lambda.base.context.LambdaContext;
import aws.lambda.base.service.S3;
import aws.lambda.base.service.SNS;
import aws.lambda.base.service.SQS;
import aws.lambda.base.service.SecretManager;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class LambdaConfig {
    static String product;
    static Context lambdaCtx; // From Lambda handleRequest context
    static LambdaConfig config;
    static List<String> secretList = new ArrayList<>();
    static List<Components> selectedComponents;

    public enum Components {
        S3, SQS, SNS
    }

    public static void init(Context ctx, String product, String[] secretList, Components[] selectedComponents)
            throws Exception {
        log.debug("Product: {}, SecretList: {}, Loaded Components: {}", product, secretList, selectedComponents);

        LambdaConfig.lambdaCtx = ctx;
        if (LambdaConfig.config != null)
            return;

        LambdaConfig.product = product;
        LambdaConfig.secretList = Arrays.asList(secretList);
        if (selectedComponents != null)
            LambdaConfig.selectedComponents = Arrays.asList(selectedComponents);

        config = new LambdaConfig();
    }

    public static LambdaConfig getInstance() {
        return config;
    }

    private LambdaContext ctx;

    private LambdaConfig() throws Exception {
        ctx = LambdaContext.getInstance();
        ctx.addToEnvironment(LambdaConstants.APPNAME_PROPERTY, LambdaConfig.product);
        loadSecrets();
        loadSelectedComponents();
        loadTruststoreFromS3ToDisk();
    }

    public void loadSecrets() throws Exception {
        String profile = ctx.getProperty("profile");
        log.info("loadSecrets, profile: {}", profile);
        if (StringUtils.isNullOrEmpty(profile))
            throw new Exception("Fail to get env profile");

        ctx.removeAllSecrets();
        StringSubstitutor sub = new StringSubstitutor(ctx.getEnvironment());

        SecretManager sm = new SecretManager();

        for (String secretTemplate : secretList) {
            String secretId = sub.replace(secretTemplate);
            log.debug("load secret: {}", secretId);
            Map<String, Object> secret = sm.fetchSecrets(secretId);
            ctx.addSecret(secret);

            if (secretId.endsWith(DbSecretConstants.API_SECRET_SUFFIX))
                ctx.setDbSecret(secret);
        }
    }

    public void loadSelectedComponents() {
        if (selectedComponents == null)
            return;

        for (Components component : selectedComponents) {
            log.info("Load {} into LambdaContext.", component);

            switch (component) {
                case S3:
                    String bucketName = ctx.getProperty("s3.bucket.name");
                    ctx.register(Components.S3.toString(), new S3(bucketName));
                    break;
                case SQS:
                    ctx.register(Components.SQS.toString(), new SQS());
                    break;
                case SNS:
                    ctx.register(Components.SNS.toString(), new SNS());
                    break;
                default:
                    log.error("Not supported component : {}", component.toString());
                    break;
            }
        }
    }

    private void loadTruststoreFromS3ToDisk() {
        String profile = ctx.getProperty("profile");
        log.info("loadFiles, profile: {}", profile);

        S3 s3 = (S3) ctx.lookup(Components.S3.toString());
        if (s3 == null)
            return;

        File file = new File(LambdaConstants.LAMBDA_TS_LOCAL_PATH); // TrustStore
        if (!file.exists())
            s3.getObject(profile + "/sslcert/app.ts", file);
    }
}
