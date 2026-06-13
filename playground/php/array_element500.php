<?php
// intialize array
$a = range(1,500);
$shuffle_times = 5;

// random it
function shuffle_array(&$a)
{
    for ($i = 0,$len = count($a); $i < $len; $i++) {
        $p = rand(0,$len-1);
        $tmp = $a[$p];
        $a[$p] = $a[$i];
        $a[$i] = $tmp;
    }
}

for ($i = 0; $i < $shuffle_times; $i++) {
    echo "shuffle $i times.\n";
    shuffle_array($a);
}

// remove a random element
$rkey = array_rand($a,1);
array_splice( $a , $rkey , 1 );
echo "After shuffled $shuffle_times times: " . json_encode($a) . "\n";


$missing = find_missing($a,count($a)+1);
echo "Missing one:".$missing."\n";

// space O(1), time O(1) , math way
function find_missing( $a , $count )
{
    $sum = ($count + 1) * $count / 2;   // Get sum of it should be
    return $sum - array_sum( $a );      // the difference between it should be and it has is the missing one.
}


?>
