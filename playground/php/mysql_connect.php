<?php

$db         = "test";

function connect_db($_db)
{
	$server     = "localhost";
	$username   = "test";
	$passwd     = "zxcvasdf";

	$conn = new mysqli($server,$username,$passwd);

	if ($conn->connect_error) {
		die("Connect to $server failed: ". $conn->connect_error);
	}
	$conn->select_db($_db);
	return $conn;
}

// get connection
$conn = connect_db($db);


try {

	// insert data to exads_test
	$sql = "insert into exads_test (`name`,`age`,`job_title`) value ('shuk4','30','swe')";
	$result = $conn->query($sql);

	
	if ($result) {
		echo "Successfully insert.\n";
	} else {
		throw new Exception("failed insert\n");
	}

	// select all datas from exads_test 
	$sql = "select `name`,`age`,`job_title` from exads_test";
	$result = $conn->query($sql);

	if ($result->num_rows > 0) {
		while ($row = $result->fetch_assoc()) {
			echo "{$row['name']} {$row['age']} {$row['job_title']}.\n";
		}
	} else {
		throw new Exception("fail or 0 result.\n");
	}

} catch (Exception $e){
	echo $e;
} finally {
	echo "connection closed.";
	$conn->close();
}



?>
