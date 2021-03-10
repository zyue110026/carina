package e2e

var deployment = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: carina-sample
  namespace: carina
  labels:
    app: sample
spec:
  replicas: 1
  selector:
    matchLabels:
      app: sample
  template:
    metadata:
      labels:
        app: sample
    spec:
      containers:
        - name: web-server
          image: docker.io/library/nginx:latest
`
