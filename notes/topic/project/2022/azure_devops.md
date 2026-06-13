# Azure DevOps

CI/CD is a key from code base to production server. With code hosting on Azure DevOps, it'll be easy to leverage Azure Pipeline to build/test/deploy. The flow is quite different than common CI/CD process. Common one is setup by build to release of one branch. Here, it is configured with build stage and release stage seperately. In earier enterprise solution, different branches in one repo have different purposes, like main project and CD configures. It is not good for gating the code quality. I convert it into single master branch way and elaborate how it configures in build and release stage. Release stage binds Build stage result and deploys to certain environments. Check out [kubernetes design] for more detail configurations.

PS: In TSMC, it has some restricted policies between on-premise and cloud. It has complex route back and forth behind the scene. When building projects, the result images will stay on the cloud and pull down to on-premise registry when doing release tasks.

### Constraints

HA and Scalability are keys for modern deployment methodology. Lack of resources causes troubles frequently in non-production environment.
