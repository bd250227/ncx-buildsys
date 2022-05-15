# ncx-buildsys
Translation of DSL content into executable code

## Workspace Justifications
This microservice is iteratively developed using `docker-compose` for simplicity and dependency minimization.  Full validation of features is done on k8s.  However, the k8s cluster is managed via the `ncx-infra` repository instead of this one.

## Notes
I would like to switch back to having all go files in a root-level `go/` directory in the project.  The `postgres` package clutters up the source directories and it won't be the last to do so.