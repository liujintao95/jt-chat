goctl model1 mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/jt_chat" -table="user,group,user_contact,contact_application"  -dir="./apps/user/model" --style=goZero

goctl model1 mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/jt_chat" -table="message,file"  -dir="./apps/message/model" --style=goZero
