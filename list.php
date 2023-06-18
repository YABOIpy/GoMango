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
                <li><a href="index.html">Tab 1</a></li>
                <li class="active">Tab 2</li>
            </ul>
        </div>
        <div class="workspace">
            <h1>Tab 2 Content</h1>
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
                        $email = $row['email'];

                        echo '<div class="form-group">';
                        echo '<label for="name">Name:</label>';
                        echo '<input type="text" id="name" name="name" value="' . $name . '" readonly>';
                        echo '</div>';

                        echo '<div class="form-group">';
                        echo '<label for="email">Email:</label>';
                        echo '<input type="email" id="email" name="email" value="' . $email . '" readonly>';
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
