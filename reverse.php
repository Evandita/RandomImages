<?php
// This script assumes it is part of the CTF challenge environment
$flag = getenv("FLAG"); // Fetch the FLAG environment variable

if ($flag) {
    echo "<script>alert('FLAG: " . htmlspecialchars($flag) . "');</script>";
} else {
    echo "<script>alert('No FLAG found in environment variables.');</script>";
}
?>
