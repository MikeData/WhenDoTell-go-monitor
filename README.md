# The "When" part of my WhenDoTell learning project

projectWhenDoTell is an attempt to build a microservice arcitecture for triggering tasks WHEN x criteria is met, to DO x process and TELL the user via x specified message.

Example:

post a task that. monitors a web page for content change, checking every 20 minutes, and sending and SMS when its finds some.

This go service is the REST endpoint for posting tasks. It then adds them to the appropriate REDIS to-do queue at the appropriate time.
