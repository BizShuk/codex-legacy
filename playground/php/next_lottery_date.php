<?php

date_default_timezone_set('Europe/Dublin');

$unix_time = time();
echo "unix time:".$unix_time."\n";
echo "w day(0 for sunday,6 for sat.):".date('w',$unix_time)."\n";
echo "current date:".date('Y-m-d H:i:s w',$unix_time)."\n";

echo getNextLotteryDate($unix_time);


//  @param int $cur , current unix time
//  @return int , next lottery date
function getNextLotteryDate($cur)
{
    /*
    $dtz = new DateTimeZone('Europe/Dublin');
    $time_in_dublin = new DateTime('now', $dtz);
    $tz_offset = $dtz->getOffset( $time_in_dublin );
    $week_sec = 7*86400;
    $diff_to_sat20 = 2*86400 + 20*3600;

    $bsat20  = ((int)(($cur - $diff_to_sat20 +$tz_offset)/($week_sec)))*($week_sec) + $diff_to_sat20 - $tz_offset + 4*86400;

    if ($bsat20 <=$cur) {
        return $bsat20 + 3*86400;
    }

    return $bsat20;
    */
    $today  = mktime(0, 0, 0, date("m",$cur)  , date("d",$cur), date("Y",$cur));
    $wday = date('w',$cur);
    $wstart = $today - $wday * 86400;

    // this week lottery dates
    $fLotteryDate = $wstart + 3 * 86400 + 20 * 3600; // week start + 3 day + 20 hour = first lottery date
    $sLotteryDate = $fLotteryDate + 3 * 86400; // first lottery date + 3 day = second lottery date

    if ( $cur >= $sLotteryDate ) {
        $fLotteryDate += 7 * 86400;
    }

    if ( $cur < $fLotteryDate) {
        return $fLotteryDate;
    } else {
        return $sLotteryDate;
    }
}








