# GPU Scheduler

GPU Scheduler is a project to help GPUs perform better on schedule. The exisiting mechanism only response to original scheduler which doesn't have any API to expose results of GPU tasks. So, the new solution plays a gatekeeper role to arrange tasks, but still using existing mechanism to launch tasks. The new solution should cover couples of aspects due to the old solution does not expose the result.

- Ability to schedule tasks
- When a task is finished, it should tell GPU Scheduler to update status of the task.It includes `SUCCESSFUL`, `FAILED`, `PROCESSING`, `SCHEDULED`.
- Common job scheduler is about to process tasks if the processor is available. In this case, it's about whether GPU is available. `The processor can not be released by only sending tasks to execute.` It has to wait the status turned into final state. `Waiting Time`
- It should have 'timeout' mechanism if the tasks is failed somewhere and does not update its status. The timeout should be based on the processing time of 90% range of specific type task.
- This application is inside of tsm organization. It's nice to have authentication, the application is sitting inside of organization. As the project grows up, it needs authorization to differentiate responsibilities between departments/projects.
- Priority Queue for sequential task process.

Tech: NodeJS, Typescript, Fastify(mvc), Jest, Redis(master-replicas), Bull(job queue framework)
Collaborative: Design and implement by myself, Peter took over when I left.

### Long-Term

The overall design for this moment is for short term. The project is in tight schedule. Leveraging existing solution is a key to achieve the goal in schedule.
It still better to control/expose overall status. This will result in incorporating the task launcher(exisiting solution part) or exposing status of tasks in task launcher.
