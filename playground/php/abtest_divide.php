<?php

//
// database: 'exads_test' (name, age, job_title , design_id[default:0])
//           'countABtest_user' (seq[int autoincrement])
// @param uname , Get from HTTP POST method param uname
// @return string , show user design group 
//

// db config
$server = "localhost";
$username = "test";
$passwd = "zxcvasdf";

// db connection
$conn = new mysqli( $server , $username , $passwd );
if ($conn->connect_error) {
    die("Connect to $server failed:". $conn->connect_error);
}

$conn->select_db("test");



// initial data
if( !isset($_GET['uname']) || empty($_GET['uname'])) {
	echo "no username";
	return;
}
$uname = $conn->real_escape_string($_GET['uname']);
$desing_id = 0;
$seq = -1;

// check user design group
$sql = "select design_id from exads_test where name='$uname'";
$result = $conn->query($sql);
if ($result->num_rows>0 ) {
    $row = $result->fetch_assoc();
    $design_id = $row['design_id'];
} else {
    echo "No such user";
    return;
}

// insert a row to countABtest_user and get the seq
if ($design_id == 0) {
    $sql = "insert into countABtest_user (seq) value ('')";
    $conn->query($sql);
    $seq = $conn->insert_id;

    if ($seq ==0) {
        echo "insert fail";
        return;
    }
    $seq = ($seq+1) % 4;
    $sql = "update exads_test set design_id=$seq where name='$uname'";
    $result = $conn->query($sql);
    
    if (!$result) {
        echo "update design_id fail.";
        return;
    }

    if ( $seq <= 1) {
        $design_id = 1;
    } elseif ( $seq <= 2 ) {
        $design_id = 2;
    } else {
        $design_id = 3;
    }
} 



echo "Do design $design_id.";





?>
