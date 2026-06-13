# gitlab ci

- project:edop/pipeline/dev-deployment/pipelines-common
  file: /execution-scope/mr/ecs-http.gitlab-ci.yml
  file: /lambda.gitlab-ci.yml
  file: container.gitlab-ci.yml
- project: edop/pipeline/standalongjobs/danger/danger-pipeline
  file: /execution-scope/edop/danger.gitlab-ci.yml

### import a remote danger file

  danger.import_dangerfile()
  