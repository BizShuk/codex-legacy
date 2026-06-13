package aws.lambda.base.context;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Hashtable;
import java.util.List;
import java.util.Map;

import aws.lambda.base.constants.ExceptionConstants;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class LambdaContext {

    private final Map<String, String> env;
    protected final Map<String, Object> bindings;
    protected final List<Map<String, Object>> secretList;
    protected Map<String, Object> dbSecret;

    // To keep it singleton and lazy loaded.
    public static class LambdaContextHelper {
        private static final LambdaContext ctx = new LambdaContext();

        private LambdaContextHelper() {
            throw ExceptionConstants.ILLEGAL_STATE_EXCEPTION;
        }
    }

    public static LambdaContext getInstance() {
        return LambdaContextHelper.ctx;
    }

    public LambdaContext() {
        log.info("Init LambdaContext...");

        Map<String, String> sysEnv = System.getenv();
        env = new Hashtable<>(sysEnv);
        log.debug("System Env: {}", env);

        bindings = new HashMap<>();
        secretList = new ArrayList<>();
    }

    public Object addToEnvironment(String propName, String propVal) {
        log.debug("addToEnvironment: {} {}", propName, propVal);
        return env.put(propName, propVal);
    }

    public Object removeToEnvironment(String propName) {
        log.debug("removeToEnvironment: {} {}", propName);
        return env.remove(propName);
    }

    public void addSecret(Map<String, Object> secret) {
        secretList.add(secret);
    }

    public void removeAllSecrets() {
        secretList.clear();
    }

    public void setDbSecret(Map<String, Object> secret) {
        dbSecret = secret;
    }

    public Map<String, Object> getDbSecret() {
        return dbSecret;
    }

    public Map<String, String> getEnvironment() {
        return env;
    }

    public void register(String beanName, Object obj) {
        bindings.put(beanName, obj);
    }

    public Object lookup(String name) {
        if (name.length() == 0)
            return this;

        Object result = bindings.get(name);
        if (result != null)
            return result;

        return null;
    }

    public String getProperty(String name) {
        String val = "";

        for (Map<String, Object> secret : secretList) {
            if (secret.containsKey(name))
                val = (String) secret.get(name);
        }

        if (env.containsKey(name))
            val = env.get(name);

        log.info("getProperty: {}, value: {}", name, val);

        return val;

    }

}
