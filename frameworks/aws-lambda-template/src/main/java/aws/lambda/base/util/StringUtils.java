package aws.lambda.base.util;

import java.nio.file.Paths;

import aws.lambda.base.constants.ExceptionConstants;

public class StringUtils {
    private StringUtils() {
        throw ExceptionConstants.ILLEGAL_STATE_EXCEPTION;
    }

    public static String maskPassword(String s) {
        return maskString(s, 2, 15);
    }

    /**
     * /** Keep reserved length of string from begin of the string and append with
     * (maxLength - s.length) of *.
     * 
     * @param s
     * @param reservedLength
     * @param maxLength
     * @return
     */
    public static String maskString(String s, int reservedLength, int maxLength) {
        if (s == null)
            return s;
        if (s.length() > reservedLength)
            s = s.substring(0, reservedLength);
        if (s.length() > maxLength) {
            s = s.substring(0, maxLength);
        } else {
            s += new String(new char[maxLength - s.length()]).replace("\0", "*");
        }
        return s;
    }

    public static String getFilenameWithoutExtension(String name) {
        if (name == null)
            return null;

        name = getFilename(name);
        int pos = name.lastIndexOf(".");
        if (pos == -1)
            return name;

        return name.substring(0, pos);
    }

    public static String getFilename(String name) {
        return Paths.get(name).getFileName().toString();
    }
}
