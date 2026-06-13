# Service Review process

1. Create MR and check pipeline pass all stages

2. Put MR and jira ticket in MR channel

3. Try to find one experience person on the particular item between Taipei members to review the MR first.

4. Get two approvals from others (1 from US at least) (k6 test task should be fine)

5. Merge MR by yourself and Gitlab CI/CD will trigger another master branch pipeline.

6. Once code is deployed to Dev, trigger pipeline stage for QA deployment.

7. Ask Stephen(backend)/Jack(frontend) to trigger cert/prod deployment for particular pipeline.

8. Add k6 integration test case, run successfully in local and send out MR (Should be fine to take this step along with 4.)

9. Once k6 integration test MR is approved. Resolved jira ticket.
