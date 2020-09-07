curl -X PATCH $1 \
  -d '{
  "updates": [
    {
      "type": "web",
      "docker_image": "'$2'"
    }
  ]
}' \
  -H "Content-Type: application/json" \
  -H "Accept: application/vnd.heroku+json; version=3.docker-releases" \
  -H "Authorization: Bearer $3"
