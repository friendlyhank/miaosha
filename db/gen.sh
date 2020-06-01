echo '0. tcp(127.0.0.1:3306)/student: 执行逆向操作'
xorm reverse -s mysql "root:123456@tcp(127.0.0.1:3306)/miaosha?charset=utf8&allowNativePasswords=true" templates/goxorm db
echo '1. 复制struct.go 到目标文件夹: xmiss_struct.go'
cp db/struct.go ./student_struct.go
echo '2. 清理零时数据'
rm -rdf ./db