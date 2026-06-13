# Recovery Strategy

## RTO, Recovery Time Objective

Duration between the event of failure and the point where operations resume

##### RTO Example

If an organization decides on an RTO of 4 hours and a disaster occurs at noon, the organization’s services need to be restored by 4PM.

## RPO

The maximum length of time permitted that data can be restored from, which may or may not mean data loss.

It's about loss tolerance, or how much data it is prepared to lose should a disaster strike.

##### RPO Example

 the same organization decides on a 5 hour RPO and a backup of data is taken at noon. The next scheduled backup is for 5PM. Any disaster that occurs between noon and 5PM can be restored to the noon backup, therefore resulting in a maximum of a 5 hour loss of data.
