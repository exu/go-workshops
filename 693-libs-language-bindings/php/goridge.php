<?php
use Spiral\Goridge;
require "vendor/autoload.php";

$rpc = new Goridge\RPC(new Goridge\SocketRelay("127.0.0.1", 6001));

// send simple string
echo $rpc->call("App.Hi", "Antony");
// send example array as map[string]string in Go
echo $rpc->call("App.Ho", ["a" => "b"]);