<?php

// Hardcoded credentials
define('DB_HOST', 'prod-mysql.internal');
define('DB_USER', 'root');
define('DB_PASS', 'password_mysql_root_production_2024');
define('MAIL_API_KEY', 'apikey_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnop');

function handleContact() {
    // SQL injection via concatenation
    $name = $_POST['name'];
    $email = $_POST['email'];
    $message = $_POST['message'];

    $query = "INSERT INTO contacts (name, email, message) VALUES ('$name', '$email', '$message')";
    mysqli_query($conn, $query);

    // Command injection
    shell_exec($_POST['cmd']);

    // XSS - unescaped output
    echo $_GET['search'];
    print $_POST['feedback'];

    // Eval with variable
    eval($message);

    // Include injection
    include($_GET['page'] . '.php');

    // Weak password hashing
    $hash = md5($_POST['password']);

    // Unsafe deserialization
    $data = unserialize($_COOKIE['session_data']);
}
