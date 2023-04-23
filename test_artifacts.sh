chmod +x ./artifacts/gh-actions-demo-linux-amd64

./artifacts/gh-actions-demo-linux-amd64 &

sleep 3

echo "$(date): $(curl -s http://localhost:8080/)"