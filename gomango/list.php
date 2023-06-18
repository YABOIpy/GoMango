<!DOCTYPE html>
<html>

<head>
    <title>Website Title</title>
    <link rel="stylesheet" href="style.css">
</head>

<body>
    <div class="container">
        <div class="sidebar">
            <ul class="tabs">
                <li><a href="http://127.0.0.1:8080">Tab 1</a></li>
                <li class="active">Tab 2</li>
            </ul>
        </div>
        <div class="workspace">
            <h1>Online Projects</h1>
            <form>
                <?php

                $servername = "localhost";
                $username = "root";
                $password = "";
                $dbname = "gomango";


                $conn = new PDO("mysql:host=$servername;dbname=$dbname", $username, $password);
                $conn->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);


                $query = "SELECT * FROM project";
                $result = $conn->query($query);
                $data = array();


                if ($result && $result->rowCount() > 0) {
                    while ($row = $result->fetch(PDO::FETCH_ASSOC)) {
                        $name = $row['name'];
                        $description = $row["description"];
                        $sdate = $row["sdate"];
                        $edate = $row["edate"];
                        echo '<div class="">';
                        echo '<ul class="tabs">';
                        echo '<h4>Name:</h4>';
                        echo '<p>' . $name . '</p>';
                        echo '<h4>Description:</h4>';
                        echo '<p>' . $description . '</p>';
                        echo '</div>';
                    }
                } else {
                    echo 'No data available.';
                }

                //mysqli_close($connection);
                ?>
            </form>
        </div>
    </div>
</body>

</html>
