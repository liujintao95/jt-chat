goctl api  go --api ./apps/user/api/user.api --dir ./apps/user/api

goctl api go --api ./apps/message/api/message.api --dir ./apps/message/api

goctl api plugin -plugin goctl-swagger="swagger -filename user-api.json -host 127.0.0.1:6001" --api ./apps/user/user.api --dir ./docs