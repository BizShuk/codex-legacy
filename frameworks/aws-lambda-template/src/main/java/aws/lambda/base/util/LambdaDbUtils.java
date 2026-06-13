package aws.lambda.base.util;

import java.nio.file.Path;
import java.nio.file.Paths;
import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.ResultSetMetaData;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import javax.sql.DataSource;

import com.zaxxer.hikari.HikariConfig;
import com.zaxxer.hikari.HikariDataSource;

import aws.lambda.base.constants.ExceptionConstants;
import aws.lambda.base.constants.LambdaConstants;
import aws.lambda.base.context.LambdaContext;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class LambdaDbUtils {

    private LambdaDbUtils() {
        throw ExceptionConstants.ILLEGAL_STATE_EXCEPTION;
    }

    private static DataSource ds;

    public static DataSource getDataSource() throws SQLException {
        if (ds != null)
            return ds;
        LambdaContext ctx = LambdaContext.getInstance();
        Map<String, Object> dbSecret = ctx.getDbSecret();
        String schema = ctx.getProperty(LambdaConstants.APPNAME_PROPERTY);
        String engine = (String) dbSecret.get("engine");
        String host = (String) dbSecret.get("host");
        String port = (String) dbSecret.get("port");
        String username = (String) dbSecret.get("username");
        String password = (String) dbSecret.get("password");
        String forceSSLprop = (String) dbSecret.get(LambdaConstants.DB_FORCESSL_PROPERTY);
        boolean forceSSL = !(forceSSLprop == null || forceSSLprop.equalsIgnoreCase("false"));
        String tsPassword = ctx.getProperty(LambdaConstants.TS_PASSWORD_PROPERTY);
        Path tsLocalPath = Paths.get(LambdaConstants.LAMBDA_TS_LOCAL_PATH);

        log.info("getDataSource: jdbc:{}//{}:{}@{}:{}/{}, forceSSL: {}, tsPassword: {}", engine, username,
                StringUtils.maskPassword(password), host, port, schema, forceSSL, StringUtils.maskPassword(tsPassword));
        HikariConfig config = new HikariConfig();
        setDriver(config, engine);
        config.setJdbcUrl(
                DbUtils.composeFullJdbcUrl(engine, host, port, schema, forceSSL, tsLocalPath.toString(), tsPassword));
        config.setUsername(username);
        config.setPassword(password);
        config.setMaximumPoolSize(1); // config.addDataSourceProperty("maximumPoolSize", 1); // Could not find
                                      // maximumPoolSize in postgresql driver

        ds = new HikariDataSource(config);
        return ds;
    }

    public static void setDriver(HikariConfig config, String enigne) {
        log.info("Setting driver for Hikari, enigne: {}", enigne);
        switch (enigne) {
            case "mysql":
                config.setDriverClassName("com.mysql.jdbc.Driver");
                break;
            case "postgres":
                config.setDriverClassName("org.postgresql.driver");
                break;
            default:
                log.info("No db engine matched in system");
                break;
        }
    }

    // @formatter:off
    public static Connection getConnection() throws SQLException {
        if (ds == null)
            ds = LambdaDbUtils.getDataSource();
        return ds.getConnection();
    }
    // @formatter:on

    public static List<Map<String, Object>> query(String sql) throws SQLException {
        log.info("query SQL: {}", sql);
        ArrayList<Map<String, Object>> result = new ArrayList<>();

        try (Connection conn = LambdaDbUtils.getConnection();
                Statement stmt = conn.createStatement();
                ResultSet rs = stmt.executeQuery(sql);) {

            ResultSetMetaData md = rs.getMetaData();
            int columns = md.getColumnCount();

            while (rs.next()) {
                HashMap<String, Object> row = new HashMap<>(columns);
                for (int i = 1; i <= columns; ++i) {
                    row.put(md.getColumnName(i), rs.getObject(i));
                }
                result.add(row);
            }
        }
        return result;
    }

    public static Map<String, Object> queryOne(String sql) throws SQLException {
        List<Map<String, Object>> results = query(sql);
        if (!results.isEmpty())
            return results.get(0);
        return null;
    }

    public static int update(String sql) throws SQLException {
        log.info("update SQL: {}", sql);
        int affectedRow = -1;
        try (Connection conn = LambdaDbUtils.getConnection(); Statement stmt = conn.createStatement();) {
            affectedRow = stmt.executeUpdate(sql);
        }
        return affectedRow;
    }

    public static int insert(String sql) throws SQLException {
        return update(sql);
    }
}
