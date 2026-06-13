package aws.lambda.base;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;

import aws.lambda.base.config.LambdaConfig;
import aws.lambda.base.config.LambdaConfig.Components;
import aws.lambda.base.context.LambdaContext;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public abstract class AbstractLambdaTemplate<E, R> implements RequestHandler<E, R> {
    protected LambdaContext ctx;

    public R handleRequest(E event, Context context) {
        try {
            LambdaConfig.init(context, getProductName(), getSecretList(), getSelectedComponentList());
            ctx = LambdaContext.getInstance();
            bindComponent(context);
            return handleService(event, context);
        } catch (Exception e) {
            log.error("Serious exception occurs in Lambda function", e);
            throw new RuntimeException(e);
        }
    }

    // It can use ctx to lookup components. Please check LambdaContext.
    protected abstract void bindComponent(Context lambdaCtx);

    protected abstract R handleService(E event, Context context) throws Exception;

    // Construct a String array with substitute key ${keyname} which will be replace
    // from env
    // variables.
    protected abstract String[] getSecretList();

    protected abstract Components[] getSelectedComponentList();

    // Use for DB schema
    protected abstract String getProductName();
}
