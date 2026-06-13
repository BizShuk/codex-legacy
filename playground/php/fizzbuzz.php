<?php

$pritned = false;
for ( $i = 1; $i <= 100; $i++ ) {
    
    if ( $i%3==0 && $printed = true ) echo "Fizz";
    if ( $i%5==0 && $printed = true ) echo "Buzz";
    if ( !$printed ) echo $i;
    $printed = false;
    echo "\n";

}

?>
