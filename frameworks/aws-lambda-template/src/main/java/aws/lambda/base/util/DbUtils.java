package aws.lambda.base.util;

import aws.lambda.base.constants.ExceptionConstants;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class DbUtils {
    private DbUtils() {
        throw ExceptionConstants.ILLEGAL_STATE_EXCEPTION;
    }

    public static final String JARPROTOCAL = "jar";
    public static final String FILEPROTOCAL = "file:";

    // load from class path inside of jar.
    public static String gettsfileURI(String tsFilePath) {
        log.debug("getTsFileURI with input tsFilePath({})", tsFilePath);
        if (tsFilePath == null)
            return null;

        String tsFileAbsPath = DbUtils.class.getClassLoader().getResource(tsFilePath).getPath();
        String protocol = (tsFileAbsPath.contains(".jar!")) ? JARPROTOCAL : FILEPROTOCAL + "//";
        tsFileAbsPath = protocol + tsFileAbsPath;
        log.info("getTsFileURI absolute path: {}", tsFileAbsPath);
        return tsFileAbsPath;
    }

    // TODO: validate db connection attributes, but with more generic way.
    public static boolean validateDbAttributes() {
        return true;
    }

    public static String composeFullJdbcUrl(String engine, String host, String port, String schema, boolean forceSSL,
            String tsFilePath, String tsPassword) {
        StringBuilder sb = new StringBuilder();
        if (engine.equals("postgres"))
            engine = "postgresql";

        sb.append(composeJdbcUrl(engine, host, port, schema));

        switch (engine) {
            case "mysql":
                sb.append("?").append(composeMysqlSSLOptions(forceSSL, tsFilePath, tsPassword));
                break;
            case "postgresql":
                sb.append("?").append(composePostgreSSLOptions(forceSSL));
                break;
            default:
                log.info("No DB engine matched in system");
                break;
        }
        return sb.toString();
    }

    public static String composeJdbcUrl(String engine, String host, String port, String schema) {
        String jdbcUrl = "jdbc:" + engine + "://" + host + ":" + port + "/" + schema;
        log.info("JDBC Url: {}", jdbcUrl);
        return jdbcUrl;
    }

    // Sample
    // "?userSSL=<boolean>&requireSSL=<boolean>&trustCertificateKeyStoreUrl=<file|classpath>:sslcert/*.ts&trustCertificateKeyPassword=<password>"
    public static String composeMysqlSSLOptions(boolean forceSSL, String tsFilePath, String tsPassword) {
        if (!forceSSL)
            return "";

        StringBuilder sb = new StringBuilder();
        sb.append("useSSL=").append(forceSSL).append("&requireSSL=").append(forceSSL)
                .append("&verifyServerCertificate=").append(forceSSL).append("&trustCertificateKeyStoreUrl=")
                .append(tsFilePath).append("&trustCertificateKeyStorePassword=").append(tsPassword).append("");

        return sb.toString();
    }

    public static String composePostgreSSLOptions(boolean forceSSL) {
        if (!forceSSL)
            return "";

        StringBuilder sb = new StringBuilder();
        sb.append("ssl=").append(forceSSL).append("&sslfactory=").append("org.postgresql.ssl.DefaultJavaSSLFactory")
                .append("&sslmode=").append("verify-ca").append("");

        // This log needs to be removed if the ssl Options include sensitive data like
        // key password
        log.info("SSL options for postgreSQL : {}", sb.toString());
        return sb.toString();
    }
}