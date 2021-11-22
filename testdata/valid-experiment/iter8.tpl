{{- $suffix := randAlphaNum 5 | lower -}}
apiVersion: batch/v1
kind: Job
metadata:
  name: experiment-{{ $suffix }}
spec:
  template:
    spec:
      containers:
      - name: iter8r
        image: kalantar/kubectl-iter8r
        imagePullPolicy: Always
        command:
        - "/bin/sh"
        - "-c"
        - |
          set -e
          # trap 'kill $(jobs -p)' EXIT

          # get experiment from secret
          sleep 5 # let secret be created
          echo getting secret experiment-{{ $suffix }}
          # kubectl get secret experiment-{{ $suffix }} -o go-template='{{"{{"}} .data.experiment {{"}}"}}' | base64 -d > experiment.yaml
          kubectl get secret experiment-{{ $suffix }} -o jsonpath='{.data.experiment}' | base64 -d | yq eval .tasks - > experiment.yaml
          
          # local run
          export LOG_LEVEL=info
          kubectl-iter8r run

          # update the secret
          kubectl create secret generic experiment-{{ $suffix }}-result --from-file=result=result.yaml --dry-run=client -o yaml | kubectl apply -f -
      restartPolicy: Never
  backoffLimit: 0
---
apiVersion: v1
kind: Secret
metadata:
  name: experiment-{{ $suffix }}
stringData:
  experiment: |
{{ . | toYAML | indent 4 }}
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: experiment-{{ $suffix }}
rules:
- apiGroups: [""]
  resources: ["secrets"]
  # resourceNames: ["experiment-{{ $suffix }}"]
  verbs: ["get", "list", "patch", "create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: experiment-{{ $suffix }}
subjects:
- kind: ServiceAccount
  name: default
roleRef:
  kind: Role
  name: experiment-{{ $suffix }}
  apiGroup: rbac.authorization.k8s.io
