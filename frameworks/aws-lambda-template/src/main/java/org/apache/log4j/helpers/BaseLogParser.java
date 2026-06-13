package org.apache.log4j.helpers;

import org.apache.log4j.spi.LoggingEvent;

import lombok.extern.slf4j.Slf4j;

@Slf4j
public class BaseLogParser extends PatternParser {
    private String profile = "L";

    private static final char PROD_CHAR = 'P';
    private static final char CERT_CHAR = 'C';
    private static final char QA_CHAR = 'Q';
    private static final char DEV_CHAR = 'D';
    private static final char JENKINS_CHAR = 'J';

    private static final char PROFILE_CHAR = 'P';
    private static final char SEVERITY_CHAR = 'N';

    public BaseLogParser(String pattern) {
        super(pattern);
        String profileProp = System.getenv("profile");

        switch (profileProp.toUpperCase().charAt(0)) {
            case PROD_CHAR:
                profile = "P";
                break;
            case CERT_CHAR:
                profile = "C";
                break;
            case QA_CHAR:
                profile = "Q";
                break;
            case DEV_CHAR:
                profile = "D";
                break;
            case JENKINS_CHAR:
                profile = "J";
                break;
            default:
                profile = "L";
                break;
        }
        log.info("profile property: {} , profile: {}", profileProp, profile);
    }

    @Override
    protected void finalizeConverter(char c) {
        PatternConverter pc = null;
        switch (c) {
            case PROFILE_CHAR:
                pc = new ProfileConverter(profile);
                currentLiteral.setLength(0);
                break;
            case SEVERITY_CHAR:
                pc = new SeverityConverter();
                currentLiteral.setLength(0);
                break;
            default:
                super.finalizeConverter(c);
                break;
        }

        if (pc != null)
            addConverter(pc);

    }

    private static class ProfileConverter extends PatternConverter {
        private String profile = "L";

        ProfileConverter(String profileEnv) {
            super();
            profile = profileEnv;
        }

        public String convert(LoggingEvent event) {
            return profile;
        }
    }

    private static class SeverityConverter extends PatternConverter {
        public String convert(LoggingEvent event) {
            return String.valueOf(event.getLevel().toInt() / 10000);
        }
    }

}
