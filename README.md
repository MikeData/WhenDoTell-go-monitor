# The "When" part of my WhenDoTell learning project.

WhenDoTell is an attempt to build a service with microservice arcitecture for triggering tasks WHEN x criteria is met, to DO x process and TELL the user via x specified means.

Example:

post a task that. monitors a web page for content change, checking every 20 minutes, and sending and SMS when its finds some.

This go service is the REST endpoint for posting tasks to. It then adds them to the appropriate REDIS to-do queue at the appropriate time (every 20 mins in our example).
