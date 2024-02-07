goctl api  go --api ./apps/user/user.api --dir ./apps/user

goctl api go --api ./apps/message/message.api --dir ./apps/message

goctl api plugin -plugin goctl-swagger="swagger -filename user-api.json -host 127.0.0.1:6001" --api ./apps/user/user.api --dir ./docs