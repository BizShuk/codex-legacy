<?php

$user = $_GET['user'];
$t = new abtest();
$design_id = $t->get_design_id($user);
if ( is_numeric($design_id) ) { 
	echo "Do design $design_id.";
}

// divide each user to one of design groups.
//
// database: 'exads_test' (name, age, job_title , design_id[default:0])
//           'countABtest_user' (seq[int autoincrement])
//
// @param user , Get from HTTP GET method param user
// @return string , show user design group 
//
class abtest{
	private $conn;
	private $server = "localhost";
	private $username = "test";
	private $passwd = "zxcvasdf";
	private $db = "test";	

	// connect db
	function connect(){
		$this->conn = new mysqli( $this->server , $this->username , $this->passwd );
		if ($this->conn->connect_error) {
			die("Connect to $this->server failed:". $this->conn->connect_error);
		}
		$this->conn->select_db("test");
	}
	
	public function get_design_id($user){
		
		if (! $this->conn) {
			$this->connect();
		}

		$uname = $this->conn->real_escape_string($user);

		$design_id = $this->check_design_id($user);
		
		if ( $design_id === false ){
			echo "No such user.";
			return false;
		} else if ( $design_id == 0 ) {
			$design_id = $this->give_design_id($user);
		}

		return $design_id;
	}

	// check user design group
	// @param string user
	// @return int design_id , false = fail
	function check_design_id($user){
		$sql = "select design_id from exads_test where name='$user'";
		$result = $this->conn->query($sql);
		if ($result->num_rows>0 ) {
			$row = $result->fetch_assoc();
			$design_id = $row['design_id'];
			return $design_id;
		} else {
			return false;
		}
	}
	

	// assign a design id to user and update to db
	// @param string user
	// @return int design_id , false = fail
	function give_design_id($user) {

		// insert a row to countABtest_user and get the seq number
		$sql = "insert into countABtest_user (seq) value ('')";
		$this->conn->query($sql);
		$seq = $this->conn->insert_id;

		if ($seq == 0 ) {
			echo "insert fail";
			return false;
		}

		// design group mapping
		$seq = ($seq+1) % 4;
		switch($seq){
			case 0:
			case 1:	$seq = 1; break;
			case 2: $seq = 2; break;
			default: $seq = 3; break;
		}

		// update design_id to user column
		$sql = "update exads_test set design_id=$seq where name='$user'";
		$result = $this->conn->query($sql);

		if (!$result) {
			echo "update design_id fail.";
			return false;
		}

		return $seq;
	}
}




?>
