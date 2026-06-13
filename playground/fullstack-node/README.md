# fullstack-practice

It's a simple project to connect every client composed by Reactjs and users can put any number to add a shared pool to sum the total number all users give. 

In this project, the diagram is like below.

Client (Reactjs) connect to server (Nodejs) and keep the connection.
Client has a form which can submit the number they give to server.
Server will queue this number in SQS and pull the message from other end point.
The end point will be responsible to sum those messages up and push the summary to all clients connecting.
The server should have region failover and the client should be able to catch up by the domain name or other ways.


The next projects selection (shold include DB failover and app failover) : 
- Rundeck (job trigger, scheduler, job componentization)
- Video streaming
- missOld
- Traveling
- PointingPoker
- 
