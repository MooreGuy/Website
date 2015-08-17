<?php

/*
	Load the url helper to help load css.
*/
$this->load->helper('url');

?>

<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title><?php echo $title;?> | Guy Moore</title>

		<!--Fontawesome -->
		<link rel="stylesheet" href="//maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css">	

		<!-- Bootstrap -->
		<link href="<? echo base_url(); ?>css/bootstrap.min.css" rel="stylesheet">
		<link href="<? echo base_url(); ?>css/personalSite.css" rel="stylesheet">	
	</head>

	<body>
		
		<div class="container">

			<div class="header">

				<nav class="navbar navbar-default">	

						<div class="navbar-header">	
							<a href="<?php echo base_url();?>" class="navbar-brand">Guy Moore</a>		
						</div>

						<ul class="nav navbar-nav pull-right">
							<li role="presentation" ><a href="<?php echo base_url();?>home/index">Info</a></li>
							<li role="Presentation" ><a href="<?php echo base_url();?>home/contact">Contact</a></li>
						</ul>		

				</nav>	
			
			</div> <!-- end header -->

			<hr>
