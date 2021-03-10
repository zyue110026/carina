package e2e

var statefulset1 = `
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: carina-stateful
  namespace: carina
spec:
  serviceName: "mysql-service"
  replicas: 2
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      terminationGracePeriodSeconds: 10
      containers:
        - name: mysqlpod
          image: mysql:5.7
          env:
            - name: MYSQL_ROOT_PASSWORD
              value: "123456"
          ports:
            - containerPort: 80
              name: my-port
          volumeMounts:
            - name: db
              mountPath: /var/lib/mysql
  volumeClaimTemplates:
    - metadata:
        name: db
      spec:
        accessModes: [ "ReadWriteOnce" ]
        storageClassName: csi-carina-sc1
        resources:
          requests:
            storage: 3Gi
`
