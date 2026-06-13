package org.apache.log4j;

import org.apache.log4j.helpers.BaseLogParser;
import org.apache.log4j.helpers.PatternParser;

public class BaseLogPatternLayout extends PatternLayout {
    @Override
    protected PatternParser createPatternParser(String pattern) {
        return new BaseLogParser(pattern);
    }
}
