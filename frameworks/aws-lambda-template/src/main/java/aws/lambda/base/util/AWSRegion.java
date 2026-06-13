package aws.lambda.base.util;

import com.amazonaws.regions.Regions;
import com.mysql.cj.util.StringUtils;
import aws.lambda.base.constants.ExceptionConstants;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class AWSRegion {
    private AWSRegion() {
        throw ExceptionConstants.ILLEGAL_STATE_EXCEPTION;
    }

    public static Regions getCurrentRegions() {
        String regionFromEnv = System.getenv("AWS_REGION");

        Regions region = null;

        if (!StringUtils.isNullOrEmpty(regionFromEnv)) {
            try {
                region = Regions.fromName(regionFromEnv);
            } catch (Exception e) {
                region = null;
            }
        }

        if (region == null) {
            try {
                region = Regions.fromName(Regions.getCurrentRegion().getName());
            } catch (Exception e) {
                region = null;
            }
        }

        if (region == null) {
            region = Regions.US_EAST_1;
        }
        log.info("Region {} is using", region);
        return region;
    }
}
