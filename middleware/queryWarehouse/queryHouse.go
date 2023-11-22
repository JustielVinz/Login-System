package querywarehouse

// Fetch data of Teacher
var GetData = `SELECT * FROM teacher`

// For Teacher query
var InsertData = `INSERT INTO teacher (name, department, created_at) VALUES (?,?,?)`

// Fetching Data for admin
var KeyGenerator = `SELECT username, password  FROM credentials WHERE username = ?`

// Inserting Data of  Student
var StudentLogin = `INSERT INTO student (id, name, student_id, department, miscellaneous, payment_method, amount, created_at) VALUES (?,?,?,?,?,?,?,?,?)`

// Creating new admin
var AdminKey = `INSERT INTO credentials (id, username, password, status, created_at) VALUES (?,?, ?, ?, ?)`
