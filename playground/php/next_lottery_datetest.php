<?php

use PHPUnit\Framework\TestCase;


/**
 * Class next_lottery_datetest
 * @author shuk
 */
class next_lottery_datetest extends TestCase
{
    /**
     * test next_lottery_date function
     *
     * @return void
     */
    public function testnext_lottery_date()
    {   
        $w0 = mktime(0,0,0,7,31,2016);
        $w7 = mktime(0,0,0,8,7,2016);
        $fl = mktime(20,0,0,8,3,2016);
        $sl = mktime(20,0,0,8,6,2016);
        $nfl= mktime(20,0,0,8,10,2016);
    
        $cases  = [];
        $expected   = [];
        
        $cases[] = $w0;     $expected[] = $fl;  // case 1
        $cases[] = $fl-1;   $expected[] = $fl;  // case 2
        $cases[] = $fl;     $expected[] = $sl;  // case 3
        $cases[] = $sl-1;   $expected[] = $sl;  // case 4
        $cases[] = $sl;     $expected[] = $nfl; // case 5
        $cases[] = $w7-1;   $expected[] = $nfl; // case 6
        $cases[] = $w7;     $expected[] = $nfl; // case 7

        echo "\n";
        for ($i = 0; $i < 7; $i++) {
            echo "case: ".date("Y-m-d H:i:s",$cases[$i]).", result: ".date("Y-m-d H:i:s",getNextLotteryDate($cases[$i]))." , expected: ".date("Y-m-d H:i:s",$expected[$i])."\n";
            $this->assertEquals($expected[$i],getNextLotteryDate($cases[$i]));
        }

        return null;
    }
    
}
